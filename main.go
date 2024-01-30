package main

import (
	"context"
	"fmt"
	"os"

	"dagger.io/dagger"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	ctx := context.Background()

	// initialize Dagger client
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	checkErr(err)
	defer client.Close()

	id, err := client.Container().From("alpine").
		Publish(ctx, "docker.io/marcosnils/test-alpine")

	fmt.Println(id, err)
}
