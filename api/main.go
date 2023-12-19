package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/serge-vm/5letters/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type SolverRequest struct {
	Ordered   map[int]string `json:"ordered" example:"3:о,5:е"`
	Unordered []string       `json:"unordered" example:"а,б"`
	Absent    []string       `json:"absent" example:"в,г"`
}

type SolverResponse []string

// @BasePath /api/v1

// QuizSolver godoc
// @Summary 5 letters quiz solver
// @Schemes
// @Description provide letter sets and get suitable variants
// @Tags solver
// @Accept json
// @Produce json
// @Param SolverRequest body SolverRequest true "Existing letters"
// @Success 200 {object} SolverResponse "List of suitable words"
// @Router /solver [post]
func QuizSolver(g *gin.Context) {
	var request SolverRequest
	g.Bind(&request)
	fmt.Println(request)
	g.JSON(http.StatusOK, []string{"your_word_1", "your_word_2"})
}

func main() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		v1.POST("/solver", QuizSolver)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8080")

}
