package result

import (
	"errors"

	"github.com/gin-gonic/gin"
)

type Result interface {
	JSON(c *gin.Context)
}

type ResultOption func(r *response)

func Err(err error) ResultOption {
	return func(r *response) {
		r.err = err
	}
}

func Data(data interface{}) ResultOption {
	return func(r *response) {
		r.data = data
	}
}

// New returns a new Result
func New(code int, msg string, options ...ResultOption) Result {
	resp := &response{
		code: code,
		err:  errors.New(msg),
	}

	for _, f := range options {
		f(resp)
	}

	return resp
}
