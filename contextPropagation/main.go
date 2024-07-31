package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "userID", 123)
	performTask(ctx)

}

func performTask(ctx context.Context) {

	userId := ctx.Value("userID")
	fmt.Println("userID :", userId)

}
