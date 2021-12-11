package middlewares

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/pretty"
	"go.uber.org/zap"

	"goadmin/pkg/logger"
	"goadmin/pkg/result"
)

var bufPool = sync.Pool{
	New: func() interface{} {
		return bytes.NewBuffer(make([]byte, 0, 2<<10)) // 2KB
	},
}

type respWriter struct {
	gin.ResponseWriter
	tee io.Writer
}

func (w *respWriter) Write(buf []byte) (int, error) {
	n, err := w.ResponseWriter.Write(buf)

	if w.tee != nil {
		w.tee.Write(buf[:n])
	}

	return n, err
}

func WrapRespWriter(c *gin.Context, tee io.Writer) {
	c.Writer = &respWriter{
		ResponseWriter: c.Writer,
		tee:            tee,
	}
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		now := time.Now().Local()

		var body []byte

		// 取出请求Body
		if (c.Request.Body != nil && c.Request.Body != http.NoBody) || !strings.Contains(c.Request.Header.Get("Content-Type"), "multipart/form-data") {
			var err error

			body, err = ioutil.ReadAll(c.Request.Body)

			if err != nil {
				result.ErrSystem(result.Err(err)).JSON(c)

				return
			}

			// 关闭原Body
			c.Request.Body.Close()

			c.Request.Body = ioutil.NopCloser(bytes.NewReader(body))
		}

		// 存储返回结果
		buf := bufPool.Get().(*bytes.Buffer)
		buf.Reset()

		defer bufPool.Put(buf)

		if strings.Contains(c.Request.Header.Get("Content-Type"), "text/html") {
			WrapRespWriter(c, buf)
		}

		c.Next()

		logger.Info(c.Request.Context(), fmt.Sprintf("[%s] %s", c.Request.Method, c.Request.URL.String()),
			zap.ByteString("params", pretty.Ugly(body)),
			zap.String("response", buf.String()),
			zap.Int("status", c.Writer.Status()),
			zap.String("duration", time.Since(now).String()),
		)
	}
}
