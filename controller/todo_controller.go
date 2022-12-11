package controller

import (
	"fmt"
	"net/http"
)

type TodoController interface {
	getTodoById(w http.ResponseWriter, r *http.Request)
	editTodoById(w http.ResponseWriter, r *http.Request)
	deleteTodoById(w http.ResponseWriter, r *http.Request)
	getTodos(w http.ResponseWriter, r *http.Request)
	createTodo(w http.ResponseWriter, r *http.Request)
}

type todoController struct {}

// 詳細
func (tc *todoController) getTodoById(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("詳細取得")
	fmt.Fprintf(w, "詳細取得")
}

func (tc *todoController) editTodoById(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("詳細取得")
	fmt.Fprintf(w, "詳細取得")
}
func (tc *todoController) deleteTodoById(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("詳細取得")
	fmt.Fprintf(w, "詳細取得")
}
// 全件
func (tc *todoController) getTodos(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("全件取得")
	fmt.Fprintf(w, "全件取得")
}
// 作成
func (tc *todoController) createTodo(w http.ResponseWriter, r *http.Request) {
	// if r.Method != http.MethodPost {
	// 	fmt.Fprintf(w, "HTTPメソッドはPOSTにしてね!!雑魚!!")
	// 	return
	// }
	fmt.Printf("作成")
	fmt.Fprintf(w, "作成")
}
