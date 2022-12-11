package handler

import (
	"fmt"
	"net/http"
)

type TodoHandler interface {
	todoByIdRequestHandler(w http.ResponseWriter, r *http.Request)
	getTodosHandler(w http.ResponseWriter, r *http.Request)
	createTodoHandler(w http.ResponseWriter, r *http.Request)
}

// 詳細
func todoByIdRequestHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case "GET":
			// controllers.getTodoById(w, r)
		case "DELETE":
			// deleteTodoById(w, r)
		case "PUT":
			// editTodoById(w, r)
		default:
				http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

// 全件取得
func getTodosHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
		fmt.Fprintf(w, "HTTPメソッドはPOSTにしてね!!雑魚!!")
		return
	}
}

// 作成
func createTodoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Fprintf(w, "HTTPメソッドはPOSTにしてね!!雑魚!!")
		return
	}
}
