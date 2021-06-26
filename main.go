package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	r := gin.Default()

	r.GET("/", func(context *gin.Context) {
		user := User{
			Name: "Tuan",
			Age:  20,
		}
		context.JSON(http.StatusOK, gin.H{
			"message": "done",
			"user":    &user,
		})
	})

	r.Run(":8080")

}
