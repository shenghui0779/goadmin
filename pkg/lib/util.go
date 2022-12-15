package lib

import (
	"crypto/rand"
	"encoding/hex"
	"goadmin/pkg/logger"
	"io"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

var CaptchaDriver base64Captcha.Driver = base64Captcha.NewDriverString(
	39,
	120,
	0,
	base64Captcha.OptionShowHollowLine,
	4,
	base64Captcha.TxtNumbers+base64Captcha.TxtAlphabet,
	nil,
	nil,
	nil,
)

func Nonce() string {
	nonce := make([]byte, 8)
	io.ReadFull(rand.Reader, nonce)

	return hex.EncodeToString(nonce)
}

func URLParamInt(c *gin.Context, key string) int64 {
	param := c.Param(key)

	if len(param) == 0 {
		return 0
	}

	v, err := strconv.ParseInt(param, 10, 64)

	if err != nil {
		logger.Err(c.Request.Context(), "err url param to int64", zap.Error(err), zap.String("param", key))

		return 0
	}

	return v
}

func URLQueryInt(c *gin.Context, key string) int64 {
	query := c.Query(key)

	if len(query) == 0 {
		return 0
	}

	v, err := strconv.ParseInt(query, 10, 64)

	if err != nil {
		logger.Err(c.Request.Context(), "err url query to int64", zap.Error(err), zap.String("query", key))

		return 0
	}

	return v
}
