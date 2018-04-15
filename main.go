package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/enjikaka/copernicus"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", copernicus.Test())
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
