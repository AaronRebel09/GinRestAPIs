package main

import (
	//"bufio"
	//"fmt"
	//"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func my_file(c *gin.Context) {
	dat, err := os.ReadFile("./my_file.txt")
	check(err)
	//fmt.Print(string(dat))

	c.IndentedJSON(http.StatusOK, string(dat))
}
