package routes

import (
	"github.com/EronAlves1996/TodoListGo/routes/db"
	"github.com/gin-gonic/gin"
)

func MarkTaskAsDone(c *gin.Context) {
	taskId := c.Query("id")

	db.MarkAsDone(taskId)

	c.Redirect(301, "/")
}
