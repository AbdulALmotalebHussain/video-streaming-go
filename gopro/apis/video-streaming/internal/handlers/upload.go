package handlers

import (
    "net/http"
    "path/filepath"
    "io/ioutil"
    "os"
    "log"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // Parse the multipart form
    err := r.ParseMultipartForm(10 << 20) // 10MB
    if err != nil {
        http.Error(w, "Error parsing form", http.StatusInternalServerError)
        return
    }

    // Retrieve the file from form data
    file, handler, err := r.FormFile("video")
    if err != nil {
        http.Error(w, "Error retrieving file", http.StatusInternalServerError)
        return
    }
    defer file.Close()

    // Create the video directory if not exists
    videoDir := "web/videos"
    if err := os.MkdirAll(videoDir, os.ModePerm); err != nil {
        log.Fatalf("Failed to create video directory: %v", err)
    }

    // Read the file content
    tempFile, err := ioutil.TempFile(videoDir, handler.Filename)
    if err != nil {
        http.Error(w, "Error creating file", http.StatusInternalServerError)
        return
    }
    defer tempFile.Close()

    fileBytes, err := ioutil.ReadAll(file)
    if err != nil {
        http.Error(w, "Error reading file", http.StatusInternalServerError)
        return
    }

    // Write the content to the temporary file
    tempFile.Write(fileBytes)

    // Redirect to home page with the video URL
    videoURL := "/videos/" + filepath.Base(tempFile.Name())
    http.Redirect(w, r, "/?video="+videoURL, http.StatusSeeOther)
}

