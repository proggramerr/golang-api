package web

import "github.com/gin-gonic/gin"

func CreateApp() *gin.Engine {
	app := gin.Default()
	InitRouters(app)
	return app
}
