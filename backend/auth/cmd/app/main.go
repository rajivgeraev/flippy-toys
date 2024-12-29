package main

import (
    "fmt"
    "net/http"
    "time"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        currentTime := time.Now().Format(time.RFC3339)
        response := fmt.Sprintf("Hello World! Current time: %s", currentTime)
        w.Write([]byte(response))
    })

    port := ":8080"
    fmt.Printf("Server is running on port %s\n", port)
    http.ListenAndServe(port, nil)
}
