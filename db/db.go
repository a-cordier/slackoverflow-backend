package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
)

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

func Connect(host string, port string, dbname string, user string, password string, sslmode string) (*gorm.DB, error) {
	params := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		host, port, dbname, user, password, sslmode,
	)
	db, err := gorm.Open("postgres", params)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&Channel{})
	db.AutoMigrate(&Question{})
	db.AutoMigrate(&Tag{})
	db.AutoMigrate(&Answer{})
	db.Model(&Answer{}).AddForeignKey("question_id", "questions(id)", "CASCADE", "CASCADE")
	db.Model(&Question{}).AddForeignKey("channel_id", "channels(id)", "CASCADE", "CASCADE")

	return db, nil
}
