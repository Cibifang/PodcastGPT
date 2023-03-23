package main

import (
    "log"
    "net/http"
    "os"

    "github.com/gorilla/mux"
    "github.com/joho/godotenv"

    "PodcastGPT/internal/api"
    "PodcastGPT/internal/config"
    "PodcastGPT/internal/repositories"
    "PodcastGPT/internal/services"
)

func main() {
    // Load environment variables from .env file
    if err := godotenv.Load(); err != nil {
        log.Print("No .env file found.")
    }

    // Create a new database connection pool
    db, err := repositories.NewDB(config.DBConfig{
        Host:     os.Getenv("DB_HOST"),
        User:     os.Getenv("DB_USER"),
        Password: os.Getenv("DB_PASSWORD"),
        DBName:   os.Getenv("DB_NAME"),
    })
    if err != nil {
        log.Fatal("Failed to connect to database: ", err)
    }
    defer db.Close()

    // Create a new chatgpt service
    gptService := services.NewChatGPTService()

    // Create a new podcast repository
    podcastRepo := repositories.NewPodcastRepository(db)

    // Create a new podcast service
    podcastService := services.NewPodcastService(podcastRepo, gptService)

    // Setup routes
    router := mux.NewRouter()
    api.SetupRoutes(router, podcastService)

    // Start HTTP Server
    srv := &http.Server{
        Addr:    "127.0.0.1:8000",
        Handler: router,
    }

    log.Print("Started server on ", srv.Addr)
    log.Fatal(srv.ListenAndServe())
}
