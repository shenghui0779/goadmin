package result

import (
	"goadmin/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type response struct {
	Code  int         `json:"code"`
	Err   bool        `json:"err"`
	Msg   string      `json:"msg"`
	ReqID string      `json:"req_id"`
	Data  interface{} `json:"data,omitempty"`
}

func (resp *response) JSON(c *gin.Context) {
	resp.ReqID = logger.GetReqID(c.Request.Context())

	c.JSON(http.StatusOK, resp)
}
