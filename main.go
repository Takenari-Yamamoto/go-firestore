package main

import (
	"fmt"
	"net/http"
)

/*
	/todo/{id}
*/
func getTodoById(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("詳細取得")
	fmt.Fprintf(w, "詳細取得")
}
func editTodoById(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("詳細取得")
	fmt.Fprintf(w, "詳細取得")
}
func deleteTodoById(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("詳細取得")
	fmt.Fprintf(w, "詳細取得")
}
/*
	/todos
*/
// TODO: リクエストに応じて処理をハンドリングする
func getTodos(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		fmt.Fprintf(w, "HTTPメソッドはGETにしてね!!雑魚!!")
		return
	}
	fmt.Printf("全件取得")
	fmt.Fprintf(w, "全件取得")
}
/*
	/todo
*/
func createTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Fprintf(w, "HTTPメソッドはPOSTにしてね!!雑魚!!")
		return
	}
	fmt.Printf("作成")
	fmt.Fprintf(w, "作成")
}

// IDによって処理を分けるためのハンドラー
func todoByIdRequestHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case "GET":
			getTodoById(w, r)
		case "DELETE":
			deleteTodoById(w, r)
		case "PUT":
			editTodoById(w, r)
		default:
				http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func main() {

	http.HandleFunc("/todo", createTodo)
 	http.HandleFunc("/todos", getTodos)
 	http.HandleFunc("/todo/{id}", todoByIdRequestHandler)

	fmt.Printf("Starting server at port 8000\n")
	http.ListenAndServe(":8000", nil)
}
