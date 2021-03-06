package ctrl

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/a-cordier/slackoverflow/core"
	"github.com/a-cordier/slackoverflow/hook"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func NewHook(engine *core.Engine) func(c *gin.Context) {
	return func(c *gin.Context) {
		var payload hook.ShortCutPayload
		logrus.Info(c.PostForm("payload"))
		if err := json.Unmarshal([]byte(c.PostForm("payload")), &payload); err != nil {
			logrus.Error("Unable to parse shortcut payload", err)
			return
		}
		if err := doSave(engine, &payload); err != nil {
			logrus.Error("Error saving payload", err)
			c.JSON(http.StatusBadRequest, hook.NewHookResponse(err.Error()))
		}
	}
}

func doSave(engine *core.Engine, payload *hook.ShortCutPayload) error {
	switch payload.CallbackID {
	case hook.AddQuestion:
		return engine.SaveQuestion(payload)
	case hook.AddAnswer:
		return engine.SaveAnswer(payload)
	default:
		return errors.New("Unknown callback ID " + payload.CallbackID)
	}
}
