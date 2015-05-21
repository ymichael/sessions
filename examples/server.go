package main

import (
	"fmt"
	"net/http"

	"github.com/ymichael/sessions"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

var (
	Sessions = sessions.NewSessionOptions("thisismysecret.", &sessions.MemoryStore{})
)

func hello(c web.C, w http.ResponseWriter, r *http.Request) {
	x := Sessions.GetSession(&c)
	if val, ok := x["count"]; ok {
		x["count"] = val.(int) + 1
	} else {
		x["count"] = 1
	}
	fmt.Fprintf(w, "Hello, %d!", x["count"])
}

func main() {
	goji.Use(Sessions.Middleware())
	goji.Get("/", hello)
	goji.Serve()
}