package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/serge-vm/5letters/internal/models"
	"github.com/serge-vm/5letters/internal/quizsolver"
)

// @BasePath /api/v1

// QuizSolver godoc
// @Summary 5 letters quiz solver
// @Schemes
// @Description provide letter sets and get suitable variants
// @Tags solver
// @Accept json
// @Produce json
// @Param SolverRequest body models.SolverRequest true "Existing letters"
// @Success 200 {object} models.SolverResponse "List of suitable words"
// @Router /solver [post]
func SolverHandler(g *gin.Context) {
	var request models.SolverRequest
	g.Bind(&request)
	g.JSON(http.StatusOK, quizsolver.Solve(request))
}
