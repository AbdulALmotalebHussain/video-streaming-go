package handlers

import (
    "html/template"
    "net/http"
    "path/filepath"
)

type HomePageData struct {
    VideoURL string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    tmplPath := filepath.Join("web", "templates", "index.html")
    tmpl, err := template.ParseFiles(tmplPath)
    if err != nil {
        http.Error(w, "Error loading template", http.StatusInternalServerError)
        return
    }

    data := HomePageData{
        VideoURL: r.URL.Query().Get("video"),
    }

    tmpl.Execute(w, data)
}

func VideoHandler(w http.ResponseWriter, r *http.Request) {
    videoName := r.URL.Path[len("/videos/"):]
    videoPath := filepath.Join("web", "videos", videoName)
    http.ServeFile(w, r, videoPath)
}

