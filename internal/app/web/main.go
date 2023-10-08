package web

import (
	"github.com/gin-gonic/gin"
	"rest_api/pkg/prometheus"
)

func CreateApp() *gin.Engine {
	app := gin.Default()
	prometheus.SetupPrometheus(app)
	InitRouters(app)
	return app
}
