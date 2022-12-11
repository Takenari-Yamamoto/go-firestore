package handler

import (
	"fmt"
	"go-firestore/controller"
	"net/http"
)

func Serve() {
	http.HandleFunc("/todo", todoCreateHandler)
 	http.HandleFunc("/todos", todosHandler)
 	http.HandleFunc("/todo/{id}", todoIdHandler)
}

// 詳細
func todoIdHandler (w http.ResponseWriter, r *http.Request) {
	var tc controller.TodoController

	switch r.Method {
		case "GET":
			tc.GetById(w, r)
		case "DELETE":
			tc.Delete(w, r)
		case "PUT":
			tc.Update(w, r)
		default:
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

// 全件取得
func todosHandler(w http.ResponseWriter, r *http.Request) {
	var tc controller.TodoController
	if r.Method != http.MethodGet {
		fmt.Fprintf(w, "HTTPメソッドはGETにしてね!!雑魚!!")
		return
	}
	tc.GetAll(w, r)
}

// 作成
func todoCreateHandler(w http.ResponseWriter, r *http.Request) {
	var tc controller.TodoController
	if r.Method != http.MethodPost {
		fmt.Fprintf(w, "HTTPメソッドはPOSTにしてね!!雑魚!!")
		return
	}
	tc.Create(w, r)
}
