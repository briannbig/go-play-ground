package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
)

const keyServerAddr = "serverAddr"

func main() {
	fmt.Println("Here we go! a simple http server in golang")

	mux := http.NewServeMux()

	mux.HandleFunc("/", getHello)
	mux.HandleFunc("/greet", greet)
	mux.HandleFunc("/hi", sayHi)
	mux.HandleFunc("/form", handleFormData)

	ctx, cancelCtx := context.WithCancel(context.Background())

	server := &http.Server{
		Addr:    ":3333",
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, keyServerAddr, l.Addr().String())
			return ctx
		},
	}

	go func() {
		err := server.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("server closed\n")
		} else if err != nil {
			fmt.Printf("error listening for server one: %s\n", err)
		}
		cancelCtx()
	}()

	<-ctx.Done()

}

func getHello(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Printf("%s got / request\n", ctx.Value(keyServerAddr))

	io.WriteString(w, "Here we go confirmed!")

}

func greet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	hasName := r.URL.Query().Has("name")
	var name string = "user"
	if hasName {
		name = r.URL.Query().Get("name")
	}
	fmt.Printf("%s got /greet request from %s\n", ctx.Value(keyServerAddr), name)
	greeting := fmt.Sprintf("Hello %s", name)
	io.WriteString(w, greeting)
}

func sayHi(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Printf("Could not read body: %s\n", err)
	}
	fmt.Printf("%s --- request body: \n %s\n", ctx.Value(keyServerAddr), body)

	io.WriteString(w, fmt.Sprintf("Hi %s", body))

}

func handleFormData(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	fmt.Printf("%s: got /form request\n", ctx.Value(keyServerAddr))

	name := r.PostFormValue("name")
	email := r.PostFormValue("email")

	if name == "" {
		name = "John"
	}
	if email == "" {
		email = "john@xyz.com"
	}

	io.WriteString(w, fmt.Sprintf("\nName: %s\nEmail: %s", name, email))

}
