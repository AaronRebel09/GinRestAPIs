package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue", Artist: "John", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughn", Artist: "Sarah", Price: 39.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)

	router.GET("/files", my_file)

	//router.GET("/counter_incr", Increment)

	//my_counter := new(Counter)

	//fmt.Println(my_counter.GetCount())
	//for i := 1; i <= 10; i++ {
	//	my_counter.Increment()
	//	}

	//	fmt.Println(my_counter.GetCount())
		
	// myCounter.count -= 6 // Defeats the encapsulation of Counter    fmt.Println(myCounter.GetCount())”


	router.Run("localhost:8080")
}
