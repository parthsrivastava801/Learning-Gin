package controllers

import (
	"go-tutorial/services"

	"github.com/gin-gonic/gin"
)

type NotesController struct {
	notesService services.NotesService
}

func (n *NotesController) InitNotesControllerRoutes(router *gin.Engine, notesService services.NotesService) {
	notes := router.Group("/notes")
	notes.GET("/", n.GetNotes())
	notes.POST("/", n.CreateNotes())
	n.notesService = notesService
}

func (n *NotesController) GetNotes() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			//"notes": "Get Request Notes",
			"notes": n.notesService.GetNotesService(),
		})
	}
}

func (n *NotesController) CreateNotes() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"notes": n.notesService.CreateNotesService(),
		})
	}
}
