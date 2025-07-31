package main

import (
    "log"
    "task_manager/data"
    "task_manager/router"
    "github.com/joho/godotenv"
    "os"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    mongoURI := os.Getenv("MONGO_URI")
    data.InitMongoDB(mongoURI)

    r := router.SetupRouter()
    r.Run(":8080")
}
