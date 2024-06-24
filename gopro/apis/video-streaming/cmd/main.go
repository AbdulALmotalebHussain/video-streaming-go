package main

import (
    "video-streaming/internal/database"
    "video-streaming/internal/cache"
    "video-streaming/internal/server"
    "os"
    "log"
    "github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    mongoURI := os.Getenv("MONGODB_URI")
	redisAddr := os.Getenv("REDIS_ADDR")
	redisPassword := os.Getenv("REDIS_PASSWORD")

    database.InitMongoDB(mongoURI)
	cache.InitRedis(redisAddr, redisPassword)

	server.StartServer()}

