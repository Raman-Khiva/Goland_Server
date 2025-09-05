package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.POST("/updateId/:id/:newId", func(ctx *gin.Context) {
		id := ctx.Param("id")
		newId := ctx.Param("newId")

		ctx.JSON(http.StatusOK, gin.H{
			"status":      "Id updated successfully",
			"old Id was ": id,
			"new Id is":   newId,
		})
	})

	log.Println("Server started successfully at port : 8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server : error", err)
	}
}
