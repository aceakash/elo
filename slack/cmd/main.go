package main

import (
	"net/http"
	"os"
	"log"
	"fmt"
	"time"
)

func main() {

	http.HandleFunc("/pool-command", func (w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.String())
		fmt.Fprintf(w, "Hello, at %s", time.Now().Format(time.RFC3339))
	});

	port := os.Getenv("PORT")

	fmt.Println("Starting on port ", port)
	err := http.ListenAndServe(":" + port, nil)
	if err != nil {
		log.Fatal("Error while starting the server", err)
	}
}
