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
