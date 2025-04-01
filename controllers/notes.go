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
	type NoteBody struct {
		Title  string `json:"title" binding:"required"`
		Status bool   `json:"status"`
	}

	return func(c *gin.Context) {
		var noteBody NoteBody
		if err := c.ShouldBindJSON(&noteBody); err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		note, err := n.notesService.CreateNotesService(noteBody.Title, noteBody.Status)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"note": note,
		})
	}

}

//c.JSON(200, gin.H{
//	"notes": n.notesService.CreateNotesService(),
//})
