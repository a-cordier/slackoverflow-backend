package core

import (
	"encoding/json"

	"github.com/a-cordier/slackoverflow/db"
	"github.com/a-cordier/slackoverflow/hook"
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
)

type Engine struct {
	db *gorm.DB
}

func NewEngine(db *gorm.DB) *Engine {
	return &Engine{db}
}

func (ng *Engine) SaveQuestion(payload *hook.ShortCutPayload) error {
	channel := mapChannel(payload)
	question := mapQuestion(payload)
	question.ChannelID = channel.ID
	return ng.db.Where(db.Channel{ID: channel.ID}).
		FirstOrCreate(channel).
		Create(question).
		Error
}

func (ng *Engine) SaveAnswer(payload *hook.ShortCutPayload) error {
	return nil
}

func mapChannel(payload *hook.ShortCutPayload) *db.Channel {
	return &db.Channel{
		ID:        payload.Channel.ID,
		Name:      payload.Channel.Name,
		Questions: []db.Question{},
	}
}

func mapQuestion(payload *hook.ShortCutPayload) *db.Question {
	q := db.NewQuestion()
	q.ID = payload.ID
	q.Text = mapText(payload.Message)
	return q
}

func mapText(message hook.Message) (jsonb postgres.Jsonb) {
	txt := extractText(message)
	b, err := json.Marshal(txt)
	if err != nil {
		return
	}
	jsonb.Scan(b)
	return
}

func extractText(message hook.Message) (elems []map[string]interface{}) {
	for _, b := range message.Blocks {
		for _, e := range b.Elements {
			for _, c := range e.Chunks {
				elems = append(elems, c)
			}
		}
	}
	return
}

func isText(chunk map[string]interface{}) bool {
	return chunk["type"] == "text"
}
