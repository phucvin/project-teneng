package main

import (
	"fmt"
    "log"
    "net/http"
    "os"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, this is a sample app, serving from %s", os.Getenv("FLY_REGION"))
    })

    log.Println("listening on", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}
