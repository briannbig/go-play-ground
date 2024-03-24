package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
)

const keyServerAddr = "serverAddr"

func main() {
	fmt.Println("Here we go! a simple http server in golang")

	mux := http.NewServeMux()

	mux.HandleFunc("/", getHello)
	mux.HandleFunc("/greet", greet)

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
