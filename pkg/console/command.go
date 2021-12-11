package console

import (
	"context"
	"goadmin/pkg/ent"
	"goadmin/pkg/ent/migrate"

	"github.com/urfave/cli/v2"
)

var Commands = []*cli.Command{
	{
		Name:    "migrate",
		Aliases: []string{"M"},
		Usage:   "数据库迁移",
		Action: func(c *cli.Context) error {
			return ent.DB.Schema.Create(context.Background(),
				migrate.WithDropIndex(true),
				migrate.WithDropColumn(true),
				migrate.WithForeignKeys(false),
			)
		},
	},
}
