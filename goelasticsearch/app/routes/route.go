package routes

import (
	"goelasticsearch/app/controllers"

	"github.com/gin-gonic/gin"
)

var GoElasticsearchController controllers.GoElasticsearchInterface = &controllers.GoElasticsearchController{}

func GetGoElasticsearch(context *gin.Context) {
	GoElasticsearchController.GetGoElasticsearch(context, context.Request)
}

func PostGoElasticsearch(context *gin.Context) {
	GoElasticsearchController.PostGoElasticsearch(context, context.Request)
}

func UpdateGoElasticsearch(context *gin.Context) {
	GoElasticsearchController.UpdateGoElasticsearch(context, context.Request)
}
