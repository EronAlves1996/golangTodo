package tasktypes

import "time"

type TaskToDb struct {
	Description string
	CreatedAt   time.Time
	ModifiedAt  time.Time
	Deadline    time.Time
	IsCompleted bool
}

type Task struct {
	Id          int32
	Description string
	CreatedAt   time.Time
	ModifiedAt  time.Time
	Deadline    time.Time
	IsCompleted bool
}

type NewTask struct {
	Description string `form:"description" binding:"required"`
	Deadline    string `form:"deadline" binding:"required"`
}

type EditTask struct {
	Id          string `form:"id" binding:"required"`
	Description string `form:"description" binding:"required"`
	Deadline    string `form:"deadline" binding:"required"`
}
