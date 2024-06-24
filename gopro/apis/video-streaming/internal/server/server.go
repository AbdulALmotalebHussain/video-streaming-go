package server

import (
    "log"
    "net/http"
    "video-streaming/internal/handlers"
)

func StartServer() {
    // Serve static files
    fs := http.FileServer(http.Dir("web/static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    // Handle routes
    http.HandleFunc("/", handlers.HomeHandler)
    http.HandleFunc("/ws", handlers.WebSocketHandler)
    http.HandleFunc("/upload", handlers.UploadHandler)
    http.HandleFunc("/videos/", handlers.VideoHandler)

    log.Println("Starting server on :8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

