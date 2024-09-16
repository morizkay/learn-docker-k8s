package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("PORT env not defined")
	}

	instanceID := os.Getenv("INSTANCE_ID")
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "the requested http method is not allowed", http.StatusMethodNotAllowed)
			return

		}

		text := "Hello World"

		if instanceID != "" {
			text = text + ". from " + instanceID
		}
		w.Write([]byte(text))
	})

	server := new(http.Server)
	server.Handler = mux
	server.Addr = ":" + port

	log.Print("Web server is starting at", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error(), "Server failed to run")
	}
}
