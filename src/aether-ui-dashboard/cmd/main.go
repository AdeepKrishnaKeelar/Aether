package main

import (
	"aether-dashboard/api"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize the Gorilla Mux Router.
	r := mux.NewRouter()

	// Register various API routes here.

	// Register the VM routes.
	api.RegisterVMRoutes(r)

	// Serve the Index page at root path.
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/dashboard/index.html")
	})

	// Serve the static files.
	fs := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))
	r.Handle("/static/", fs)

	// Start the server.
	fmt.Println("Starting the server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))

}
