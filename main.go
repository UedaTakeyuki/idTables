package main

import (
	//	"log"

	"github.com/gin-gonic/gin"
)

const DataFileJson = "idTable.json"
const DataFileCSV = "idTable.csv"

func main() {

	// idTable, idCount
	//	idCount, _ := NewPersistentInt("idTableCount")
	//	idTable, _ := NewIDTable("idTable", idCount)

	// routes
	r := gin.Default()
	//r.Static("/assets", "./assets")
	//	r.Use(static.Serve("/", static.LocalFile("./assets", false)))
	//	r.LoadHTMLGlob("templates/*.html")

	v1 := r.Group("/v1")

	// idTableName, internalID, externalID
	v1.POST("/add", func(c *gin.Context) {
		v1AddHandler(c)
	})

	// idTableName, internalID, externalID
	v1.POST("/update", func(c *gin.Context) {
		v1AddHandler(c)
	})

	/*
		v1.GET("/list/:idTableName", func(c *gin.Context) {
			v1ListHandler(c, &(idTable.table))
		})

		v1.GET("/list_invert/:idTableName", func(c *gin.Context) {
			v1ListHandler(c, &(idTable.table))
		})
	*/
	// idTableName, externalID
	v1.POST("/delete_by_externalid", func(c *gin.Context) {
		v1DeleteByExternalIDHandler(c)
	})

	/*
		// idTableName, internalID
		v1.POST("/delete_by_internalid/", func(c *gin.Context) {
			v1DeleteHandler(c, idTable)
		})

		v1.GET("/internalid/:idTableName", func(c *gin.Context) {
			externalID := c.Param("id")
			c.JSON(200, gin.H{
				"uedaname": idTable.table[externalID],
			})
		})

		v1.GET("/externalid/:idTableName", func(c *gin.Context) {
			v1LabelInfoHandler(c, idTable)
		})

		v1.POST("/updateRedirectURL", func(c *gin.Context) {
			v1UpdateRedirectURL(c)
		})
	*/
	r.Run(":10002") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	//log.Fatal(autotls.Run(r, "isolde.uedasoft.com", "ml.uedasoft.com"))
}
