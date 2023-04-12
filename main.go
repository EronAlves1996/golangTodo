package main

import (
	"database/sql"
	"os"

	"github.com/EronAlves1996/TodoListGo/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	initializeDb()
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	addRoutes(r)
	r.Run(":8080")
}

func addRoutes(e *gin.Engine) {
	e.Static("/static", "./static")
	e.GET("/", routes.Index)
}

func initializeDb() {
	file, err := os.Open("./db/tasks.db")

	if err != nil {
		err := os.Mkdir("db", os.ModePerm)

		if err != nil {
			panic(err)
		}

		db, err := sql.Open("sqlite3", "./db/tasks.db")

		if err != nil {
			panic(err)
		}

		sqlStmt := `CREATE TABLE tasks (
			id INTEGER NOT NULL AUTOINCREMENT PRIMARY KEY,
			description varchar(50) not null,
			created_at timestamp not null,
			modified_at timestamp not null,
			deadline timestamp not null,
			is_completed boolean not null
		);`

		if _, err := db.Exec(sqlStmt); err != nil {
			panic(err)
		}
	}

	file.Close()
}
