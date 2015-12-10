package handlers

import (
	"net/http"

	"github.com/ThatsMrTalbot/example"
)

// Index handler
type Index struct {
	*example.Mux `inject:"private"`

	HelloHandler *Hello `inject:""`
}

// Start handler
func (index *Index) Start() error {
	index.HandleFunc("/", index.Index)
	index.SubRoute("/hello", index.HelloHandler)

	return nil
}

// Index handler
func (index *Index) Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Index"))
}
