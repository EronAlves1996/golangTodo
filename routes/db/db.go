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
		tasks = append(tasks, mapRow(rows))
	}

	return tasks
}

func MarkAsDone(id string) {
	db := attemptOpenDb()
	defer db.Close()

	sqlStmt := `UPDATE tasks
		SET is_completed=true,
		modified_at=?
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

	if _, err := stmt.Exec(time.Now(), id); err != nil {
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

func GetTask(id string) tasktypes.Task {
	db := attemptOpenDb()
	defer db.Close()

	stmt, err := db.Prepare("select * from tasks where id=?")

	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	rst, err := stmt.Query(id)

	if err != nil {
		panic(err)
	}

	defer rst.Close()

	rst.Next()
	task := mapRow(rst)
	return task
}

func SaveTask(t *tasktypes.EditTask) {
	db := attemptOpenDb()
	defer db.Close()

	tx, err := db.Begin()

	if err != nil {
		panic(err)
	}

	stmt, err := tx.Prepare(`UPDATE tasks
	SET description=?,
	deadline=?,
	modified_at=?
	WHERE id=?`)

	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	tm, err := time.Parse("2006-1-02", t.Deadline)

	if err != nil {
		panic(err)
	}

	if _, err := stmt.Exec(t.Description, tm, time.Now(), t.Id); err != nil {
		panic(err)
	}

	tx.Commit()
}

func DeleteTask(id string) {
	db := attemptOpenDb()
	defer db.Close()

	tx, err := db.Begin()

	if err != nil {
		panic(err)
	}

	stmt, err := tx.Prepare("DELETE FROM tasks WHERE id=?")

	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	if _, err := stmt.Exec(id); err != nil {
		panic(err)
	}

	tx.Commit()
}

func mapRow(r *sql.Rows) tasktypes.Task {
	var id int32
	var description string
	var created_at time.Time
	var modified_at time.Time
	var deadline time.Time
	var is_completed bool
	err := r.Scan(&id, &description, &created_at, &modified_at, &deadline, &is_completed)

	if err != nil {
		panic(err)
	}

	return tasktypes.Task{
		Id:          id,
		Description: description,
		CreatedAt:   created_at,
		ModifiedAt:  modified_at,
		Deadline:    deadline,
		IsCompleted: is_completed,
	}
}
