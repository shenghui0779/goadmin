package result

import (
	"fmt"
	"goadmin/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type response struct {
	code int
	err  error
	data interface{}
}

func (resp *response) JSON(c *gin.Context) {
	obj := gin.H{
		"code": resp.code,
		"err":  false,
		"msg":  fmt.Sprintf("[%s] %s", logger.GetReqID(c.Request.Context()), resp.err.Error()),
	}

	if resp.code != 0 {
		obj["err"] = true
	}

	if resp.data != nil {
		obj["data"] = resp.data
	}

	c.JSON(http.StatusOK, obj)
}
