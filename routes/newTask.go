package routes

import (
	"time"

	"github.com/EronAlves1996/TodoListGo/routes/db"
	"github.com/EronAlves1996/TodoListGo/routes/tasktypes"
	"github.com/gin-gonic/gin"
)

type NewTask struct {
	Description string `form:"description" binding:"required"`
	Deadline    string `form:"deadline" binding:"required"`
}

func CreateTask(c *gin.Context) {
	c.HTML(200, "new-task.tmpl", gin.H{
		"test": "Test data",
	})
}

func HandleNewTask(c *gin.Context) {
	newTask := NewTask{}

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
