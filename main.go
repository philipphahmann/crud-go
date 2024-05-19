package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/philipphahmann/crud-go/src/controller/routes"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Erro loading .env file")
	}

	router := gin.Default()

	routes.Init(&router.RouterGroup)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
