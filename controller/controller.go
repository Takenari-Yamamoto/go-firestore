package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type TodoController struct {}

type TodoItem struct {
	Id uuid.UUID `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}

type TodoRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

var todoList []TodoItem

func (tc TodoController) Create(w http.ResponseWriter, r *http.Request) {

	// 参考：https://blog.logrocket.com/how-to-make-http-post-request-with-json-body-in-go/
	reqBody, _ := ioutil.ReadAll(r.Body)
	var todo TodoRequest
	json.Unmarshal(reqBody, &todo)
	json.NewEncoder(w).Encode(todo)
	fmt.Println(string(todo.Title))

	todoList = append(todoList, TodoItem{
		Id: uuid.New(),
		Title: todo.Title,
		Content: todo.Content,
		CreatedAt: time.Now(),
	})

}

func (tc TodoController) GetById(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("詳細取得")
	fmt.Fprintf(w, "詳細取得")
}

func (tc TodoController) GetAll(w http.ResponseWriter, r *http.Request) {

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(&todoList); err != nil {
		log.Fatal(err)
	}
	_, err := fmt.Fprint(w, buf.String())
	if err != nil {
		return
	}
}

func (tc TodoController) Update(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("編集")
	fmt.Fprintf(w, "編集")
}

func (tc TodoController) Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("削除")
	fmt.Fprintf(w, "削除")
}
