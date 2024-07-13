package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "<h1>Full Cycle Rocks!!</h1>")
	})
	http.ListenAndServe(":1000", nil)
}
