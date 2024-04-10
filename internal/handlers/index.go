package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/serge-vm/5letters/assets"
)

func IndexHandler(g *gin.Context) {
	index, _ := assets.HTML.ReadFile("html/index.html")
	g.Data(http.StatusOK, "text/html; charset=utf-8", index)
}
