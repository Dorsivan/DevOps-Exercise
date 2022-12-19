package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

type song struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

var songs = []song{
    {ID: "1", Title: "Enter Sandman", Artist: "Metallica", Price: 19.99},
    {ID: "2", Title: "Castle", Artist: "Halsey", Price: 19.99},
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}

func getSongs(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, songs)
}

func getSong(c *gin.Context) {
    for _, element := range songs {
        if element.ID == c.Param("id") {
            c.IndentedJSON(200, element)
            return
        }
    }
    c.String(404, "--404-- Song not found")
}

func main() {
    router := gin.Default()
    router.GET("/albums", getAlbums)
    router.GET("/songs", getSongs)
    router.GET("/song/:id", getSong)
    router.Run("0.0.0.0:8080")
}