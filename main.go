package main

import (
	"crypto/sha256"
	"net/http"
	"strings"
)

func main() {
	// endpoints
	http.HandleFunc("/health/", health)
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}

func health(w http.ResponseWriter, r *http.Request) {

	hash := sha256.New()
	message := strings.TrimPrefix(r.URL.Path, "/health/")
	sha := hash.Sum([]byte(message))
	w.Write(sha)
}
