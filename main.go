package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Here we go! a simple http server in golang")

	mux := http.NewServeMux()

	mux.HandleFunc("/", getHello)
	mux.HandleFunc("/greet", greet)

	err := http.ListenAndServe(":3333", mux)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("Server closed")
	} else if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}

}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("request body: %d\n", r.Body)
	io.WriteString(w, "Here we go confirmed!")

}

func greet(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello user")
}
