package main

import (
	"fmt"
	"net/http"

	"github.com/ymichael/sessions"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

var (
	// Sessions = sessions.NewSessionOptions(
	// 	"thisismysecret.",
	// 	&sessions.MemoryStore{})
	Sessions = sessions.NewSessionOptions(
		"thisismysecret.",
		sessions.NewRedisStore("tcp", "localhost:6379"))
)

func hello(c web.C, w http.ResponseWriter, r *http.Request) {
	x := Sessions.GetSessionObject(&c)
	fmt.Println(x)
	if val, ok := x["count"]; ok {
		x["count"] = val.(int) + 1
	} else {
		x["count"] = 1
	}
	fmt.Fprintf(w, "Hello, %d!", x["count"])
}

func destroy(c web.C, w http.ResponseWriter, r *http.Request) {
	Sessions.RegenerateSession(&c, w)
	http.Redirect(w, r, "/", 302)
}

func main() {
	goji.Use(Sessions.Middleware())
	goji.Get("/", hello)
	goji.Get("/destroy", destroy)
	goji.Serve()
}
