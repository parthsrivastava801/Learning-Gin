package main

import (
	//"net/http"
	"go-tutorial/controllers"
	internal "go-tutorial/internal/database"
	"go-tutorial/services"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	db := internal.InitDB()

	if db == nil {
		//error while connecting to the databse
	}

	notesService := &services.NotesService{}
	notesService.InitService(db)

	notesController := &controllers.NotesController{}
	notesController.InitNotesControllerRoutes(router, *notesService) //mujhe lund nahi pata asterk kyu

	router.Run(":8000")

}
