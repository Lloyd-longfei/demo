package router

import (
	"demo/handler"
	middleLog "demo/middleware/log"

	"github.com/gin-contrib/cors"
	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func RouteLink() *gin.Engine {

	middleLog.GenerateOrOutPutLogger()

	router := gin.Default()

	router.Use(cors.Default())

	gin.SetMode(gin.DebugMode)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/", func(c *gin.Context) {
		log.Error("hello world")
	})
	router.POST("/insert", handler.Test)
	return router
}
