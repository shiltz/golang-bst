package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type album struct {
	ID string `json:id`
}

var albums = []album{{ID: "1"}}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {

	setupServer().Run("localhost:8080")
}

func setupServer() *gin.Engine {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	return router
}
