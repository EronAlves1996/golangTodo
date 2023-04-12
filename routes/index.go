package routes

import (
	"github.com/EronAlves1996/TodoListGo/routes/db"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {

	tasks := db.GetTasks()

	c.HTML(200, "index.tmpl", gin.H{
		"tasks": tasks,
	})

}
