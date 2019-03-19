package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/iamwalala/romannumerals"
)

func romanNumberHandler(w http.ResponseWriter, r *http.Request) {
	urlPathElements := strings.Split(r.URL.Path, "/")
	// If request is GET with correct syntax
	if urlPathElements[1] == "roman_number" {
		number, _ := strconv.Atoi(strings.TrimSpace(urlPathElements[2]))
		if number == 0 || number > 10 {
			// If resource is not in the list, send Not Found status
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - Not Found"))
		} else {
			fmt.Fprintf(w, "%q\n", html.EscapeString(romannumerals.Numerals[number]))
		}
	} else {
		// For all other requests, tell that Client sent a bad request
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Bad request"))
	}
}

func main() {
	http.HandleFunc("/", romanNumberHandler)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
