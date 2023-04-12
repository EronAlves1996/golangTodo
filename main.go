package main

import (
	"github.com/EronAlves1996/TodoListGo/routes"
	"github.com/EronAlves1996/TodoListGo/routes/db"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db.InitializeDb()

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	addRoutes(r)

	r.Run(":8080")
}

func addRoutes(e *gin.Engine) {
	e.Static("/static", "./static")
	e.GET("/", routes.Index)
	e.GET("/nova-tarefa", routes.CreateTask)
	e.POST("/criar-tarefa", routes.HandleNewTask)
	e.GET("/concluir-tarefa", routes.MarkTaskAsDone)
	e.GET("/editar-tarefa", routes.EditTask)
	e.POST("/salvar-tarefa", routes.SaveTask)
	e.GET("/deletar-tarefa", routes.DeleteTask)
}
