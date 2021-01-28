package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	Init()
	route := gin.Default()

	route.GET("/role", getRole)
	route.GET("/role/:newid", getRole)

	route.Run(":8080")
}

func getRole(c *gin.Context) {
	newid := c.Param("newid")
	if newid != "" {
		if newid, err := strconv.Atoi(newid); err == nil {
			c.JSON(200, Data[newid-1])
		} else {
			c.String(400, "[Bad Request]: {{newid}}} is invalid")
		}

		return
	}

	c.JSON(200, Data)
}
