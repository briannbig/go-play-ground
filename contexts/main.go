package main

import (
	"context"
	"fmt"
)

const keyUsername = "username"

func main() {
	ctx := context.Background()

	ctx = context.WithValue(ctx, keyUsername, "jdoe")
	printUsername(ctx)
}

func printUsername(ctx context.Context) {

	fmt.Printf("initial %s's value is %s\n><><><><><><cleaning username...\n", keyUsername, ctx.Value(keyUsername))

	ctx = cleanUsername(ctx)

	fmt.Printf("Cleaned %s's value is %s\n", keyUsername, ctx.Value(keyUsername))

}

func cleanUsername(ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, keyUsername, fmt.Sprintf("~%s", ctx.Value(keyUsername)))
	fmt.Printf("cleaning %s's value to %s\n", keyUsername, ctx.Value(keyUsername))

	return ctx

}
