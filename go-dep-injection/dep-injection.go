package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// PrintGreeting prints greeting on the given io writer
func PrintGreeting(writer io.Writer, greeting string) {
	fmt.Fprintf(writer, greeting)
}

func myGreetHandler(w http.ResponseWriter, r *http.Request) {
	PrintGreeting(w, "Helloooo!!")
}

func main() {
	PrintGreeting(os.Stdout, "Hello, world")
	http.ListenAndServe(":5000", http.HandlerFunc(myGreetHandler))
}
