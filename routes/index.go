package routes

import (
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {

	c.HTML(200, "index.tpml", gin.H{
		"test": "Test data",
	})

}
