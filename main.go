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
	route.DELETE("/role/:newid", deleteRole)

	route.Run(":8080")
}

func getRole(c *gin.Context) {
	newid := c.Param("newid")
	if newid != "" {
		if newid, err := strconv.Atoi(newid); err == nil {
			unewid := uint(newid)
			for _, role := range Data {
				if role.ID == unewid {
					c.JSON(200, role)
				}
			}
		} else {
			c.String(400, "Item ID is invalid.")
		}

		return
	}

	c.JSON(200, Data)
}

func deleteRole(c *gin.Context) {
	newid := c.Param("newid")
	if newid, err := strconv.Atoi(newid); err == nil {
		unewid := uint(newid)
		for index, role := range Data {
			if role.ID == unewid {
				copy(Data[index:], Data[index+1:])
				Data[len(Data)-1] = Role{}
				Data = Data[:len(Data)-1]

				c.String(200, "Item is deleted successfully.")
				return
			}
		}

		c.String(404, "Item is not found.")
	} else {
		c.String(400, "Item ID is invalid.")
	}
}
