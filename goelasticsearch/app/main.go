package main

import (
	"goelasticsearch/app/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/module/apmgin"
)

func Setup() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Authorization", "Content-Type", "Access-Control-Allow-Origin"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "HEAD"}
	router.Use(cors.New(config))
	router.Use(apmgin.Middleware(router))

	router.GET("/search", routes.GetGoElasticsearch)
	router.POST("/insert", routes.PostGoElasticsearch)
	router.POST("/update", routes.UpdateGoElasticsearch)

	return router
}

func main() {
	router := Setup()
	router.Run(":80")
}
