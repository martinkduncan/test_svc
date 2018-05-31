package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// gitSHA from build
var gitSHA string

type HealthResponse struct {
	GitSHA string `json:"gitSHA,omitempty"`
}

func main() {
	// endpoints
	http.HandleFunc("/health/", healthHandler(gitSHA, nil))
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}

func healthHandler(gitSHA string, f func() error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("In closure")
		// throw 500 if there is nothing
		if f != nil {
			if err := f(); err != nil {
				http.Error(w, "", http.StatusInternalServerError)
			}
		}

		fmt.Println("gitSHA: " + gitSHA)
		// Do the work
		h := HealthResponse{
			GitSHA: gitSHA,
		}
		b, err := json.Marshal(h)
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
		}
		w.Write(b)
	}
}
