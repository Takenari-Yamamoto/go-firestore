package main

import (
	"fmt"
	"go-firestore/handler"
	"net/http"
)

func main() {

	fmt.Printf("Starting server at port 8000\n")
	handler.Serve()
	http.ListenAndServe(":8000", nil)
}
