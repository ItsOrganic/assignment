package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/", Handler)
	r.Run(":8080")
}

func Handler(c *gin.Context) {
	var numbers []int32
	if err := c.BindJSON(&numbers); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json body"})
		return
	}
	var res int32

	for _, num := range numbers {
		res += num
	}
	c.JSON(http.StatusOK, gin.H{"result": res})
}
