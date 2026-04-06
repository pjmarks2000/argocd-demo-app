package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "This is Latchu! Hello from ArgoCD + GitHub Actions + GitOps!")
    })

    fmt.Println("Server running on port 80")
    http.ListenAndServe(":80", nil)
}
