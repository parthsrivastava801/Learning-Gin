package services

import (
	"fmt"
	internal "go-tutorial/internal/models"

	"gorm.io/gorm"
)

type NotesService struct {
	db *gorm.DB
}

func (n *NotesService) InitService(database *gorm.DB) {
	n.db = database
	n.db.AutoMigrate(&internal.Notes{})
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

func (n *NotesService) CreateNotesService() Note {
	data := Note{
		Id: 3, Name: "Note 3",
	}
	err := n.db.Create(&internal.Notes{
		Id:     1,
		Title:  "Notes",
		Status: false,
	})
	if err != nil {
		fmt.Print(err)
	}
	return data
}
