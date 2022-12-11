package controller

import (
	"fmt"
	"net/http"
)

type TodoController struct {}

func (tc TodoController) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("作成")
	fmt.Fprintf(w, "作成")
}

func (tc TodoController) GetById(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("詳細取得")
	fmt.Fprintf(w, "詳細取得")
}

func (tc TodoController) GetAll(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("全件取得")
	fmt.Fprintf(w, "全件取得")
}

func (tc TodoController) Update(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("編集")
	fmt.Fprintf(w, "編集")
}

func (tc TodoController) Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("削除")
	fmt.Fprintf(w, "削除")
}
