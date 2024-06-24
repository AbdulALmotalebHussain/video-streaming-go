package handlers

import (
    "github.com/gorilla/websocket"
    "net/http"
    "log"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println("Failed to upgrade connection:", err)
        http.Error(w, "Failed to upgrade connection", http.StatusInternalServerError)
        return
    }
    defer conn.Close()

    // Example: Echo received messages
    for {
        messageType, message, err := conn.ReadMessage()
        if err != nil {
            log.Println("Read error:", err)
            break
        }
        err = conn.WriteMessage(messageType, message)
        if err != nil {
            log.Println("Write error:", err)
            break
        }
    }
}

