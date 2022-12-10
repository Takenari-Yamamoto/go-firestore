package main

import (
	"net/http"
	"fmt"
)

func main() {
	fmt.Printf("Starting server at port 8080\n")
	http.ListenAndServe(":8000", nil)
}
