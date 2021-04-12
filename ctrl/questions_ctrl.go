package ctrl

import (
	"net/http"

	"github.com/a-cordier/slackoverflow/api"
	"github.com/a-cordier/slackoverflow/core"
	"github.com/gin-gonic/gin"
)

func NewQuestionListHandler(engine *core.Engine) func(*gin.Context) {
	return func(ctx *gin.Context) {
		questions, err := engine.FindQuestions()
		if err != nil {
			api.WriteError(ctx, err, http.StatusInternalServerError)
			return
		}
		ctx.JSON(http.StatusOK, questions)
	}
}
