package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type BookMetaData struct {
	Title   string `json:"title"`
	Creator string `json:"creator"`
}

func main() {
	fmt.Printf("hello, world\n")

	route := gin.Default()
	route.PUT("/book", createBook)
	route.Run()
}

func createBook(c *gin.Context) {
	// var book BookMetaData

	data := new(BookMetaData)

	err := c.BindJSON(data)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	c.String(200, fmt.Sprintf("%#v", data))

}
