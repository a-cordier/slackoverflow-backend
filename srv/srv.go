package srv

import (
	"time"

	"github.com/a-cordier/slackoverflow/core"
	"github.com/a-cordier/slackoverflow/ctrl"
	"github.com/a-cordier/slackoverflow/db"
	"github.com/gin-gonic/contrib/ginrus"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func StartServer(port string) {
	db, err := db.Connect("localhost", "5432", "slackoverflow", "slackoverflow", "slackoverflow", "disable")
	if err != nil {
		logrus.Panicln("Cannot connect to database")
	}
	defer db.Close()

	engine := core.NewEngine(db)
	router := gin.New()
	router.Use(ginrus.Ginrus(logrus.StandardLogger(), time.RFC3339, true))
	router.POST("/hook", ctrl.NewHook(engine))
	router.GET("/questions", ctrl.NewQuestionListHandler(engine))
	router.GET("/answers/:questionID", ctrl.NewAnswerListHandler(engine))
	router.Run(":" + port)
}
