package core

import (
	"encoding/json"

	"github.com/a-cordier/slackoverflow/api"
	"github.com/a-cordier/slackoverflow/db"
	"github.com/a-cordier/slackoverflow/hook"
	"github.com/jinzhu/gorm/dialects/postgres"
)

type Engine struct {
	db *db.DB
}

func NewEngine(db *db.DB) *Engine {
	return &Engine{db}
}

func (ng *Engine) FindQuestions() ([]api.Question, error) {
	dbqs, err := ng.db.FindQuestions()
	if err != nil {
		return nil, err
	}
	apiqs := make([]api.Question, len(dbqs))
	for i, q := range dbqs {
		apiq, err := toApiQuestion(&q)
		if err != nil {
			continue
		}
		apiqs[i] = *apiq
	}
	return apiqs, nil
}

func (ng *Engine) FindAnswers(qID string) ([]api.Answer, error) {
	dbas, err := ng.db.FindAnswers(qID)
	if err != nil {
		return nil, err
	}
	apias := make([]api.Answer, len(dbas))
	for i, a := range dbas {
		apia, err := toApiAnswer(&a)
		if err != nil {
			continue
		}
		apias[i] = *apia
	}
	return apias, nil
}

func (ng *Engine) SaveQuestion(payload *hook.ShortCutPayload) error {
	channel := mapChannel(payload)
	question := mapQuestion(payload)
	return ng.db.SaveQuestion(channel, question)
}

func (ng *Engine) SaveAnswer(payload *hook.ShortCutPayload) error {
	question, err := ng.db.FindQuestion(payload.ThreadID)
	if err != nil {
		return err
	}
	answer := mapAnswer(payload)
	answer.QuestionID = question.ID
	return ng.db.SaveAnswer(answer)
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
	q.ID = getID(payload)
	q.Text = mapText(payload.Message)
	return q
}

func toApiQuestion(dbq *db.Question) (*api.Question, error) {
	apiq := &api.Question{}
	apiq.ID = dbq.ID
	txt, err := dbq.Text.Value()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(txt.([]byte), &apiq.Text)
	if err != nil {
		return nil, err
	}
	return apiq, nil
}

func toApiAnswer(dba *db.Answer) (*api.Answer, error) {
	apia := &api.Answer{}
	apia.ID = dba.ID
	apia.QuestionID = dba.QuestionID
	txt, err := dba.Text.Value()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(txt.([]byte), &apia.Text)
	if err != nil {
		return nil, err
	}
	return apia, nil
}

func mapAnswer(payload *hook.ShortCutPayload) *db.Answer {
	a := db.NewAnswer()
	a.ID = payload.ID
	a.Text = mapText(payload.Message)
	return a
}

func getID(payload *hook.ShortCutPayload) string {
	if "" != payload.ThreadID {
		return payload.ThreadID
	}
	return payload.ID
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
