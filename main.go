package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/a-h/templ"
)

const (
	PORT = ":8080"
)

type GlobalState struct {
	Count int
}

var global GlobalState

func getHandler(w http.ResponseWriter, r *http.Request) {
	component := page(global.Count, 0)
	component.Render(r.Context(), w)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if r.Form.Has("global") {
		global.Count++
	}

	getHandler(w, r)
}

type NowHandler struct {
	Now func() time.Time
}

func NewNowHandler(now func() time.Time) NowHandler {
	return NowHandler{Now: now}
}

func (nh NowHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	timeComponent(nh.Now()).Render(r.Context(), w)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			postHandler(w, r)
			return
		}
		getHandler(w, r)
	})
	http.Handle("/time", NewNowHandler(time.Now))
	http.Handle("/404", templ.Handler(notFoundComponent(), templ.WithStatus(http.StatusNotFound)))

	fmt.Printf("listenning on port %s\n", PORT)
	http.ListenAndServe(PORT, nil)
}
