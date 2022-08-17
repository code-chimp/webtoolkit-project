package main

import (
	"github.com/code-chimp/webtoolkit"
	"log"
	"net/http"
)

func main() {
	mux := routes()

	log.Println("Starting server on port 8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func routes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("."))))
	mux.HandleFunc("/download", downloadFile)

	return mux
}

func downloadFile(w http.ResponseWriter, r *http.Request) {
	var t webtoolkit.Tools

	t.DownloadStaticFile(w, r, "./files", "tipfinger.jpg", "loljohnny.jpg")
}
