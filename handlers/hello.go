package handlers

import (
	"net/http"

	"github.com/ThatsMrTalbot/example"
	"github.com/ThatsMrTalbot/example/queues"
)

// Hello handler
type Hello struct {
	*example.Mux `inject:"private"`

	Queues *example.Queues `inject:""`
}

// Start handler
func (hello *Hello) Start() error {
	hello.GetFunc("/", hello.Index)
	hello.GetFunc("/:name", hello.Hello)
	hello.GetFunc("/mail", hello.Mail)
	hello.PostFunc("/mail", hello.PostMail)

	return nil
}

// Index handler
func (hello *Hello) Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

// Hello handler
func (hello *Hello) Hello(w http.ResponseWriter, r *http.Request) {
	name := example.GetValue(r, "name")
	w.Write([]byte("Hello " + name))
}

// Mail handler
func (hello *Hello) Mail(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<form method=\"POST\">Email:<input name=\"email\" type=\"text\" /><button>Go</button></form>"))
}

// PostMail handler
func (hello *Hello) PostMail(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")

	job := &queues.MailJob{
		Email: email,
	}
	hello.Queues.MailJobQueue.Put(job)

	w.Write([]byte("Check your email"))
}
