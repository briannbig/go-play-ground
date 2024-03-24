package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Here we go!")
	http.HandleFunc("/", getHello)

	err := http.ListenAndServe(":3333", nil)

	if err != nil {
		fmt.Println("Error starting http server", err)
	} else {
		fmt.Println("Server started")
	}
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("request body: %d\n", r.Body)
	io.WriteString(w, "Here we go confirmed!")

}
