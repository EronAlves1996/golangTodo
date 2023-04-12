package main

import (
	"github.com/EronAlves1996/TodoListGo/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Run(":8080")
}

func addRoutes(e *gin.Engine) {
	e.GET("/", routes.Index)
}
