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

func isText(chunk map[string]interface{}) bool {
	return chunk["type"] == "text"
}

func isEmoji(chunk map[string]interface{}) bool {
	return chunk["type"] == "emoji"
}

func mapTags(message hook.Message) (tags []db.Tag) {
	for _, b := range message.Blocks {
		for _, e := range b.Elements {
			for _, c := range e.Chunks {
				if isEmoji(c) {
					tags = append(tags, db.Tag{ID: c["name"].(string)})
				}
			}
		}
	}
	return tags
}

func extractText(message hook.Message) (elems []map[string]interface{}) {
	for _, b := range message.Blocks {
		for _, e := range b.Elements {
			for _, c := range e.Chunks {
				if isText(c) {
					elems = append(elems, c)
				}
			}
		}
	}
	return
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

func mapQuestion(payload *hook.ShortCutPayload) *db.Question {
	q := db.NewQuestion()
	q.ID = payload.ID
	q.Tags = mapTags(payload.Message)
	q.Text = mapText(payload.Message)
	return q
}

func mapChannel(payload *hook.ShortCutPayload) *db.Channel {
	return &db.Channel{
		ID:        payload.Channel.ID,
		Name:      payload.Channel.Name,
		Questions: []db.Question{},
	}
}

func (ng *Engine) SaveQuestion(payload *hook.ShortCutPayload) error {
	channel := mapChannel(payload)
	question := mapQuestion(payload)
	question.ChannelID = channel.ID
	ng.db.Where(db.Channel{ID: channel.ID}).FirstOrCreate(channel)
	ng.db.Create(question)
	return nil
}

func (ng *Engine) SaveAnswer(payload *hook.ShortCutPayload) error {
	return nil
}
