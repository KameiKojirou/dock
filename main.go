package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// spaHandler serves static files and falls back to index.html when the file isn't found.
type spaHandler struct {
	staticPath string // e.g., "frontend/dist"
	indexPath  string // e.g., "index.html"
}

func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Get the absolute path to the requested file
	path := filepath.Join(h.staticPath, r.URL.Path)

	// Check whether the file exists
	info, err := os.Stat(path)
	if err != nil || info.IsDir() {
		// File doesn't exist or is a folder, so serve the index.html
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	}

	// Otherwise, serve the file using the standard file server.
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}

func main() {
	// Create our SPA file server handler
	spa := spaHandler{staticPath: "frontend/dist", indexPath: "index.html"}

	// Set the root route to be handled by our SPA handler
	http.Handle("/", spa)

	// Start the server on port 1323
	log.Println("Serving Nuxt app on http://localhost:1323")
	if err := http.ListenAndServe(":1323", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
