package main

import (
	"fmt"
	"net/http"
)

const version = "1.0.0"

func main() {
	err := http.ListenAndServe("0.0.0.0:8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(fmt.Sprintf("I'm version '%s'", version)))
		if err != nil {
			panic(err)
		}
	}))
	if err != nil {
		panic(err)
	}
}
