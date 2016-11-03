package main

import (
	"net/http"
	"os"
	"log"
	"fmt"
	"time"
	"strings"
	"github.com/gorilla/handlers"
)

func main() {
	r := http.NewServeMux()

	RespondToPoolCommands := func (w http.ResponseWriter, r *http.Request) {
		text := strings.ToLower(r.URL.Query().Get("text"))

		fmt.Fprintf(w, "[%s] %s", time.Now().Format(time.RFC3339), text)
	}

	// Only log requests to our admin dashboard to stdout
	r.Handle("/pool-command", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(RespondToPoolCommands)))

	port := os.Getenv("PORT")

	fmt.Println("Starting on port ", port)
	err := http.ListenAndServe(":" + port, handlers.CompressHandler(r))
	if err != nil {
		log.Fatal("Error while starting the server", err)
	}
}
