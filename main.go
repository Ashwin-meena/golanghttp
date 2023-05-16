package main

import (
    "fmt"
    "log"
	"go-dummy/database"
	"go-dummy/router"
    "github.com/joho/godotenv"
	"net/http"
)

func main() {
    loadEnv()
    loadDatabase()
    serveApplication()
}

func loadEnv() {
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatal("Error loading .env file")
    }
}

func loadDatabase() {
    database.Connect()
}

func serveApplication() {
    r := router.Router()
	r.Use()
    fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}