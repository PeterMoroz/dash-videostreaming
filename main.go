package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	
		
	http.HandleFunc("/", ShowVideos)
	http.ListenAndServe(":" + port, nil)
}

func ShowVideos(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("videos")
	if err != nil {
		log.Println("Could not open directory ('videos'). ", err.Error())
		http.Error(w, "could not list available videos", http.StatusInternalServerError)
		return
	}
	
	items, err := f.Readdir(0)
	if err != nil {
		log.Println("Could not read directory ('videos') content. ", err.Error())
		http.Error(w, "could not list available videos", http.StatusInternalServerError)
		return
	}
	
	var videosDirs []string	
	for _, item := range items {
		if item.IsDir() {
			videosDirs = append(videosDirs, item.Name())			
		}
	}
	
	fp := path.Join("templates", "index.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		log.Println("Could not parse html-template. ", err.Error())
		http.Error(w, "could not render html-page", http.StatusInternalServerError)
		return
	}
	
	if err := tmpl.Execute(w, videosDirs); err != nil {
		log.Println("Could not execxute html-template. ", err.Error())
		http.Error(w, "could not render html-page", http.StatusInternalServerError)
	}
}