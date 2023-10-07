package web

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"rest_api/app/album"
	docs "rest_api/docs"
)

func InitRouters(engine *gin.Engine) {
	docs.SwaggerInfo.BasePath = ""

	album.InitAlbumRouter(engine)
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
