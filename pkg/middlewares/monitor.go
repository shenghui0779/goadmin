package middlewares

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	httpRequestCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "demo",
		Subsystem: "api",
		Name:      "requests_count",
		Help:      "The total number of http request",
	}, []string{"method", "path", "status"})

	httpRequestDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "demo",
		Subsystem: "api",
		Name:      "duration_seconds",
		Help:      "The http request latency in seconds",
	}, []string{"method", "path", "status"})
)

func init() {
	prometheus.MustRegister(httpRequestCounter)
	prometheus.MustRegister(httpRequestDuration)
}

// Monitor 监控请求次数，时长
func Monitor() gin.HandlerFunc {
	return func(c *gin.Context) {
		begin := time.Now().Local()

		httpRequestCounter.With(prometheus.Labels{
			"method": c.Request.Method,
			"path":   c.Request.URL.Path,
			"status": strconv.Itoa(c.Writer.Status()),
		}).Inc()

		c.Next()

		httpRequestDuration.With(prometheus.Labels{
			"method": c.Request.Method,
			"path":   c.Request.URL.Path,
			"status": strconv.Itoa(c.Writer.Status()),
		}).Observe(time.Since(begin).Seconds())
	}
}
