package main

import (
	"github.com/gin-gonic/gin"
	"github.com/serge-vm/5letters/api/handlers"
	"github.com/serge-vm/5letters/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		v1.POST("/solver", handlers.SolverHandler)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, ginSwagger.DefaultModelsExpandDepth(-1)))
	r.Run(":8080")

}
