package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	var counter int = 0
	env := os.Getenv("color")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		counter++
		fmt.Fprintf(w, fmt.Sprintf("%s Counter Value: %d\n", env, counter))
	})

	http.ListenAndServe(":8080", nil)
}
