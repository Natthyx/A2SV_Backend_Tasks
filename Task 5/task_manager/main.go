package main

import (
    "task_manager/data"
    "task_manager/router"
)

func main() {
    mongoURI := "mongodb+srv://a2sv:password123$@cluster0.tlowqek.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
    data.InitMongoDB(mongoURI)

    r := router.SetupRouter()
    r.Run(":8080")
}
