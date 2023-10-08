package prometheus

import (
	"github.com/Depado/ginprom"
	"github.com/gin-gonic/gin"
)

func SetupPrometheus(engine *gin.Engine) {
	//registry := prometheus.NewRegistry()
	p := ginprom.New(
		ginprom.Engine(engine),
		ginprom.Subsystem("gin"),
		ginprom.Path("/metrics"),
	)

	engine.Use(p.Instrument())
}
