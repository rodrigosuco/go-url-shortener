package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rodrigosuco/go-url-shortener/internal/database"
	"github.com/rodrigosuco/go-url-shortener/routes"
)

func main() {

	err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

	database.DbConnect()

	r:= gin.Default()

	routes.SetupRoutes(r)

	r.Run(":3000")
}