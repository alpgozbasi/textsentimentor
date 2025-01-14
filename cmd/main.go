package main

import (
	"log"

	"github.com/alpgozbasi/textsentimentor/internal/api"
	"github.com/gin-gonic/gin"
)

func main() {
	// create gin engine
	r := gin.Default()

	// setup routes
	api.SetupRoutes(r)

	// run server
	log.Fatal(r.Run(":8080"))

}
