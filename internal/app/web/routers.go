package web

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"rest_api/api/docs"
	albumHandlers "rest_api/internal/app/handlers/album"
)

func InitRouters(engine *gin.Engine) {
	docs.SwaggerInfo.BasePath = ""
	albumHandlers.InitAlbumRouter(engine)
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
