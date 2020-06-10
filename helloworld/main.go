package main

import (
	"fmt"
	"net/http"
)

//PrintHello retrieves all the tasks depending on the
func PrintHello(name string) string {
	return "Hello, you've requested " + name
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, PrintHello("adam"))
	})

	http.ListenAndServe(":8081", nil)
}
