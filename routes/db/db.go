package db

import (
	"database/sql"
	"os"
	"time"

	"github.com/EronAlves1996/TodoListGo/routes/tasktypes"
)

func InitializeDb() {
	file, err := os.Open("./routes/db/tasks.db")

	if err != nil {
		db := attemptOpenDb()
		defer db.Close()

		sqlStmt := "CREATE TABLE tasks (" +
			"id INTEGER PRIMARY KEY," +
			"description varchar(50) not null," +
			"created_at timestamp not null," +
			"modified_at timestamp not null," +
			"deadline timestamp not null," +
			"is_completed boolean not null);"

		if _, err := db.Exec(sqlStmt); err != nil {
			panic(err)
		}

	}

	file.Close()
}

func GetTasks() []tasktypes.Task {
	db := attemptOpenDb()
	defer db.Close()

	sqlStmt := "SELECT * FROM tasks;"

	rows, err := db.Query(sqlStmt)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	tasks := make([]tasktypes.Task, 0)

	for rows.Next() {
		var id int32
		var description string
		var created_at time.Time
		var modified_at time.Time
		var deadline time.Time
		var is_completed bool
		err := rows.Scan(&id, &description, &created_at, &modified_at, &deadline, &is_completed)

		if err != nil {
			panic(err)
		}

		tasks = append(tasks, tasktypes.Task{
			Id:          id,
			Description: description,
			CreatedAt:   created_at,
			ModifiedAt:  modified_at,
			Deadline:    deadline,
			IsCompleted: is_completed,
		})
	}

	return tasks
}

func MarkAsDone(id string) {
	db := attemptOpenDb()
	defer db.Close()

	sqlStmt := `UPDATE tasks
		SET is_completed=true
		WHERE id=?`

	tx, err := db.Begin()

	if err != nil {
		panic(err)
	}

	stmt, err := tx.Prepare(sqlStmt)

	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	if _, err := stmt.Exec(id); err != nil {
		panic(err)
	}

	tx.Commit()
}

func CreateTask(t *tasktypes.TaskToDb) {
	db := attemptOpenDb()
	defer db.Close()

	sqlStmt := `INSERT INTO tasks(
		description, 
		created_at, 
		modified_at, 
		deadline, 
		is_completed)
		values (?, ?, ?, ?, ?)`

	tx, err := db.Begin()

	if err != nil {
		panic(err)
	}

	stmt, err := tx.Prepare(sqlStmt)

	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	if _, err := stmt.Exec(t.Description, t.CreatedAt, t.ModifiedAt, t.Deadline, t.IsCompleted); err != nil {
		panic(err)
	}

	tx.Commit()
}

func attemptOpenDb() *sql.DB {
	db, err := sql.Open("sqlite3", "./routes/db/tasks.db")

	if err != nil {
		panic(err)
	}

	return db
}
