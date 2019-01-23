package router

import (
	"demo/handler"
	middleLog "demo/middleware/log"

	"github.com/gin-contrib/cors"
	log "github.com/sirupsen/logrus"

	_ "demo/docs"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title Swagger Example
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v1
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
