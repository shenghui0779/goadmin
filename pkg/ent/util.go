package ent

import (
	"context"
	"database/sql"
	"fmt"
	"goadmin/pkg/logger"
	"runtime/debug"

	"entgo.io/ent/dialect"
	"github.com/shenghui0779/yiigo"
	"go.uber.org/zap"
)

// DB ent client.
var DB *Client

func InitDB() {
	DB = NewClient(Driver(dialect.DebugWithContext(yiigo.EntDriver(), func(ctx context.Context, v ...interface{}) {
		logger.Info(ctx, "SQL info", zap.String("SQL", fmt.Sprint(v...)))
	})))
}

type TxHandler func(ctx context.Context, tx *Tx) error

// Transaction Executes ent transaction with callback function.
// The provided context is used until the transaction is committed or rolledback.
func Transaction(ctx context.Context, callback TxHandler) error {
	tx, err := DB.Tx(ctx)

	if err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			logger.Err(ctx, "ent transaction handler panic", zap.Any("error", r), zap.ByteString("stack", debug.Stack()))

			rollback(ctx, tx)
		}
	}()

	if err = callback(ctx, tx); err != nil {
		rollback(ctx, tx)

		return err
	}

	if err = tx.Commit(); err != nil {
		rollback(ctx, tx)

		return err
	}

	return nil
}

func rollback(ctx context.Context, tx *Tx) {
	if err := tx.Rollback(); err != nil && err != sql.ErrTxDone {
		logger.Err(ctx, "err ent transaction rollback", zap.Error(err))
	}
}
