package result

import "github.com/gin-gonic/gin"

const CodeOK = 0

type Result interface {
	JSON(c *gin.Context)
}

type ResultOption func(r *response)

func Err(err error) ResultOption {
	return func(r *response) {
		r.Msg = err.Error()
	}
}

func Data(data interface{}) ResultOption {
	return func(r *response) {
		r.Data = data
	}
}

// New returns a new Result
func New(code int, msg string, options ...ResultOption) Result {
	resp := &response{
		Code: code,
		Msg:  msg,
	}

	if code != CodeOK {
		resp.Err = true
	}

	for _, f := range options {
		f(resp)
	}

	return resp
}
