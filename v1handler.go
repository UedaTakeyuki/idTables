package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// idTables
var idTables = make(map[string]*IDTable)

func getIdTable(idTableName string) (idTable *IDTable, err error) {
	// create idTable if not exist?
	err = nil
	_, isExist := idTables[idTableName]
	if isExist {
		idTable = idTables[idTableName]
	} else {
		idTable, err = NewIDTable(idTableName)
		if err == nil {
			idTables[idTableName] = idTable
		}
	}
	return
}

func pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func v1AddHandler(c *gin.Context) {
	// parse Form
	c.Request.ParseForm()
	fmt.Println(c.Request.Form)
	idTableName := c.Request.Form["idTableName"][0]
	internalID := c.Request.Form["internalID"][0]
	externalID := c.Request.Form["externalID"][0]
	fmt.Println(c.Request.Form["idTableName"])
	fmt.Println(c.Request.Form["internalID"])
	fmt.Println(c.Request.Form["externalID"])

	idTable, err := getIdTable(idTableName)
	if err != nil {
		c.JSON(500, gin.H{
			"command": "add",
			"result":  "error",
			"message": err.Error(),
		})
	}

	err = idTable.Update(internalID, externalID)
	if err != nil {
		c.JSON(500, gin.H{
			"command": "add",
			"result":  "error",
			"message": err.Error(),
		})
		c.Abort()
	} else {
		c.JSON(200, gin.H{
			"command": "add",
			"result":  "ok",
		})
	}
}

func v1DeleteByExternalIDHandler(c *gin.Context) {
	// parse Form
	c.Request.ParseForm()
	fmt.Println(c.Request.Form)
	idTableName := c.Request.Form["idTableName"][0]
	//	internalID  = c.Request.Form["internalID"]
	externalID := c.Request.Form["externalID"][0]
	//	fmt.Println(c.Request.Form["internalID"])

	idTable, err := getIdTable(idTableName)
	if err != nil {
		c.JSON(500, gin.H{
			"command": "add",
			"result":  "error",
			"message": err.Error(),
		})
	}

	idTable.DeleteDeleteByExternalID(externalID)
	c.JSON(200, gin.H{
		"command": "delete_by_externalid",
		"result":  "ok",
	})
}
