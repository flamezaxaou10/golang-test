package gingo

import (
	"gingo/beer"

	_ "gingo/docs"

	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	//_ "github.com/swaggo/gin-swagger/example/basic/docs"
)

func StartServer() {
	r := gin.New()
	r.Use(gin.Recovery())
	url := ginSwagger.URL("http://localhost:1234/docs/doc.json")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	// ===== Prometheus
	// get global Monitor object
	m := ginmetrics.GetMonitor()
	// +optional set metric path, default /debug/metrics
	m.SetMetricPath("/metrics")
	// +optional set slow time, default 5s
	m.SetSlowTime(10)
	// +optional set request duration, default {0.1, 0.3, 1.2, 5, 10}
	// used to p95, p99
	m.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})
	// set middleware for gin
	m.Use(r)

	beer.NewRoutes(r)

	r.Run(":1234") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
