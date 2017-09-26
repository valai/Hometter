package main

import (
	"github.com/gin-gonic/gin"

	"hometter/controller"
//	"hometter/db"
)

func main() {
	router := gin.Default()

	router.GET("/users", controller.SelectUsers)

	router.Run(":8000")
}
