package routes

import (
	"time"

	"github.com/EronAlves1996/TodoListGo/routes/db"
	"github.com/EronAlves1996/TodoListGo/routes/tasktypes"
	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	c.HTML(200, "new-task.tmpl", gin.H{
		"Title":       "Nova Tarefa",
		"Destination": "/criar-tarefa",
		"ButtonLabel": "Criar",
	})
}

func HandleNewTask(c *gin.Context) {
	newTask := tasktypes.NewTask{}

	if err := c.ShouldBind(&newTask); err != nil {
		panic(err)
	}

	tm, err := time.Parse("2006-01-02", newTask.Deadline)

	if err != nil {
		panic(err)
	}

	taskToSave := tasktypes.TaskToDb{
		Description: newTask.Description,
		CreatedAt:   time.Now(),
		ModifiedAt:  time.Now(),
		Deadline:    tm,
		IsCompleted: false,
	}

	db.CreateTask(&taskToSave)

	c.Redirect(301, "/")
}
