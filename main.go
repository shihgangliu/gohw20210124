package main

import (
	"encoding/json"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	Init()
	gin.ForceConsoleColor()
	route := gin.Default()

	route.GET("/role", getRole)
	route.POST("/role", createRole)
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

func createRole(c *gin.Context) {
	var newItem Role
	decoder := json.NewDecoder(c.Request.Body)
	if err := decoder.Decode(&newItem); err == nil {
		skillIndex := uint(1)
		for i := range newItem.Skills {
			newItem.Skills[i].ID = skillIndex
			skillIndex++
		}

		newItem.ID = uint(len(Data)) + 1
		Data = append(Data, newItem)

		c.String(200, "New item is created successfully.")
		return
	}

	c.String(400, "New item is created unsuccessfully.")
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
