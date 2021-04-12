package ctrl

import (
	"net/http"

	"github.com/a-cordier/slackoverflow/api"
	"github.com/a-cordier/slackoverflow/core"
	"github.com/gin-gonic/gin"
)

func NewAnswerListHandler(engine *core.Engine) func(*gin.Context) {
	return func(ctx *gin.Context) {
		answers, err := engine.FindAnswers(ctx.Param("questionID"))
		if err != nil {
			api.WriteError(ctx, err, http.StatusInternalServerError)
			return
		}
		ctx.JSON(http.StatusOK, answers)
	}
}
