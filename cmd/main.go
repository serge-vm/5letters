package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/serge-vm/5letters/assets"
	"github.com/serge-vm/5letters/docs"
	"github.com/serge-vm/5letters/internal/config"
	"github.com/serge-vm/5letters/internal/handlers"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	config := config.NewAPIConfig()

	r := gin.Default()

	r.SetTrustedProxies([]string{"127.0.0.1"})

	r.StaticFS(config.BaseUrl+"/static", http.FS(assets.Static))

	docs.SwaggerInfo.BasePath = config.BaseUrl + "/api/v1"

	r.GET(config.BaseUrl+"/", handlers.IndexHandler)
	v1 := r.Group(config.BaseUrl + "/api/v1")
	{
		v1.POST("/solver", handlers.SolverHandler)
	}
	r.GET(config.BaseUrl+"/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, ginSwagger.DefaultModelsExpandDepth(-1)))
	r.Run(config.Host + ":" + strconv.Itoa(config.Port))

}
