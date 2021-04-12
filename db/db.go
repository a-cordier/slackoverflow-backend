package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
)

type DB struct {
	orm *gorm.DB
}

type Channel struct {
	ID        string `gorm:"primaryKey"`
	Name      string `gorm:"uniqueIndex"`
	Questions []Question
}

type Question struct {
	ID        string `gorm:"primaryKey"`
	Text      postgres.Jsonb
	Tags      []Tag `gorm:"many2many:question_tags;"`
	Answers   []Answer
	ChannelID string `gorm:"uniqueIndex"`
}

type Tag struct {
	ID string `gorm:"primaryKey"`
}

type Answer struct {
	ID         string `gorm:"primaryKey"`
	Text       postgres.Jsonb
	QuestionID string `gorm:"uniqueIndex"`
}

func NewQuestion() *Question {
	return &Question{
		ID:        "",
		Text:      postgres.Jsonb{},
		Tags:      []Tag{},
		Answers:   []Answer{},
		ChannelID: "",
	}
}

func NewAnswer() *Answer {
	return &Answer{
		ID:         "",
		Text:       postgres.Jsonb{},
		QuestionID: "",
	}
}

func Connect(host string, port string, dbname string, user string, password string, sslmode string) (*DB, error) {
	params := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		host, port, dbname, user, password, sslmode,
	)
	db, err := gorm.Open("postgres", params)
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&Channel{}).
		AutoMigrate(&Question{}).
		AutoMigrate(&Tag{}).
		AutoMigrate(&Answer{}).
		Model(&Answer{}).AddForeignKey("question_id", "questions(id)", "CASCADE", "CASCADE").
		Model(&Question{}).AddForeignKey("channel_id", "channels(id)", "CASCADE", "CASCADE").
		Error

	return &DB{db}, err
}

func (db *DB) Close() {
	db.orm.Close()
}

func (db *DB) FindQuestions() ([]Question, error) {
	var questions []Question
	err := db.orm.Find(&questions).Error
	return questions, err
}

func (db *DB) FindQuestion(ID string) (*Question, error) {
	var question Question
	err := db.orm.Find(&question, ID).Error
	return &question, err
}

func (db *DB) SaveQuestion(channel *Channel, question *Question) error {
	question.ChannelID = channel.ID
	return db.orm.Where(Channel{ID: channel.ID}).
		FirstOrCreate(channel).
		Create(question).
		Error
}

func (db *DB) FindAnswers(qID string) ([]Answer, error) {
	var answers []Answer
	err := db.orm.Where("question_id = ?", qID).Find(&answers).Error
	return answers, err
}

func (db *DB) SaveAnswer(answer *Answer) error {
	return db.orm.Create(answer).Error
}
