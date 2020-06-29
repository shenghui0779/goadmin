package helpers

import (
	"strconv"

	"github.com/shenghui0779/yiigo"
	"go.uber.org/zap"
)

// Int 字符串转int
func Int(s string) int {
	n, err := strconv.Atoi(s)

	if err != nil {
		yiigo.Logger().Error("convert string to int error", zap.String("error", err.Error()))
	}

	return n
}

// Int64 字符串转int64
func Int64(s string) int64 {
	n, err := strconv.ParseInt(s, 10, 64)

	if err != nil {
		yiigo.Logger().Error("convert string to int64 error", zap.String("error", err.Error()))
	}

	return n
}

// Float64 字符串转float64
func Float64(s string) float64 {
	n, err := strconv.ParseFloat(s, 64)

	if err != nil {
		yiigo.Logger().Error("convert string to float64 error", zap.String("error", err.Error()))
	}

	return n
}
