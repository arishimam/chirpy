package main

import (
	// "fmt"
	"log"
	"net/http"
	// "path"
)

func main() {
	const port = ":8080"
	const filepath = "."

	mux := http.NewServeMux()
	mux.Handle("/app/", http.StripPrefix("/app/", http.FileServer(http.Dir(filepath))))

	mux.HandleFunc("/healthz", readinessHandler)

	server := http.Server{
		Handler: mux,
		Addr:    port,
	}

	log.Printf("Serving files from %s on port %s\n", filepath, port)
	log.Fatalf("%w", server.ListenAndServe())
}

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))

}
