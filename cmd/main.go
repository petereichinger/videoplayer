package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/petereichinger/videoplayer/templates"
)

func main() {
	component := templates.Hello("John")
	http.Handle("/", templ.Handler(component))

	fmt.Println("Listening on :42069")
	http.ListenAndServe(":42069", nil)
}
