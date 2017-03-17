package main

import (
	"fmt"
	"github.com/719ben/liar/liar"
	"net/http"
)

func humanHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "You are a human, nice.")
}

func botHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "You are a bot, nice.")
}

func main() {
	middleware := liar.New(http.HandlerFunc(humanHandler), http.HandlerFunc(botHandler))
	http.Handle("/", middleware)
	http.ListenAndServe(":8000", nil)
}
