package routes

import (
	"github.com/EronAlves1996/TodoListGo/routes/db"
	"github.com/EronAlves1996/TodoListGo/routes/tasktypes"
	"github.com/gin-gonic/gin"
)

func MarkTaskAsDone(c *gin.Context) {
	taskId := c.Query("id")

	db.MarkAsDone(taskId)

	c.Redirect(301, "/")
}

func EditTask(c *gin.Context) {
	taskId := c.Query("id")

	task := db.GetTask(taskId)

	c.HTML(200, "new-task.tmpl", gin.H{
		"Title":           "Editar Tarefa",
		"Destination":     "/salvar-tarefa",
		"TaskDescription": task.Description,
		"TaskDeadline":    task.Deadline.Format("2006-01-02"),
		"TaskId":          taskId,
		"ButtonLabel":     "Salvar",
	})
}

func SaveTask(c *gin.Context) {
	taskToSave := tasktypes.EditTask{}

	err := c.Bind(&taskToSave)

	if err != nil {
		panic(err)
	}

	db.SaveTask(&taskToSave)

	c.Redirect(301, "/")
}
