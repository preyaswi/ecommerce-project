package main

import (
	"fmt"
	"log"
	"net/http"
	"newone/infrastructure"
	"newone/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	_, err := infrastructure.ConnectToDB()
	if err != nil {
		log.Fatal(err)
	}
	router := gin.Default()
	routes.UserRouter(router)

	fmt.Println("Starting server on port 8080...")
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
