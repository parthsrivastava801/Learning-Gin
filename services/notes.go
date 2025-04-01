package services

import (
	"errors"
	"fmt"
	internal "go-tutorial/internal/models"

	"gorm.io/gorm"
)

type NotesService struct {
	db *gorm.DB
}

func (n *NotesService) InitService(database *gorm.DB) {
	n.db = database
	//n.db.AutoMigrate(&internal.Notes{}) discovering an alternative way
}

type Note struct {
	Id   int
	Name string
}

func (n *NotesService) GetNotesService() []Note {
	data := []Note{
		{Id: 1, Name: "Note 1"},
		{Id: 2, Name: "Note 2"},
	}
	return data
}

func (n *NotesService) CreateNotesService(title string, status bool) (*internal.Notes, error) {
	note := &internal.Notes{
		Title:  title,
		Status: status,
	}

	if note.Title == "" {
		return nil, errors.New("title is required")
	}

	if err := n.db.Create(note).Error; err != nil {
		fmt.Print(err)
		return nil, err
	}
	return note, nil
}
