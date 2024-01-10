package main

import (
	"context"
	"flag"
	"fmt"
	"math"
	"math/rand"
	"os"

	"dagger.io/dagger"
)

func main() {
	ctx := context.Background()
	env := flag.String("env", "dev", "environment for dagger")
	flag.Parse()

	// initialize Dagger client
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stderr))
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// use a golang:1.21 container
	// mount the source code directory on the host
	// at /app in the container
	source := client.Container().
		From("golang:1.21").
		WithEnvVariable("CGO_ENABLED", "0").
		WithDirectory("/app", client.Host().Directory(".", dagger.HostDirectoryOpts{
			Exclude: []string{"ci/", ".devcontainer/"},
		}))

	// set the working directory in the container
	// install application dependencies
	runner := source.WithWorkdir("/app").
		WithExec([]string{"go", "mod", "tidy"})

	// run application tests
	test := runner.WithExec([]string{"go", "test", "./src/"})

	// build application
	buildDir := test.WithExec([]string{"go", "build", "-o", "/build/ping-socket", "./src/"}).
		Directory("/build")

	// use an alpine:3.19 container
	// copy the build/ directory from the build stage
	// add binary as entrypoint
	// publish the resulting container to a registry
	container := client.Container().
		From("alpine:3.19").
		WithDirectory("/", buildDir).
		WithEntrypoint([]string{"/ping-socket"})

	if *env == "dev" {
		exportPath := fmt.Sprintf("%s/ping-socket.tar", ".")

		_, err := container.Export(ctx, exportPath)
		if err != nil {
			panic(err)
		}

		fmt.Printf("DEV environment: exported file to %s \nTo run use docker load -i %s \n", exportPath, exportPath)
		return
	}

	ref, err := container.Publish(ctx, fmt.Sprintf("ttl.sh/ping-socket-%.0f", math.Floor(rand.Float64()*10000000))) //#nosec
	if err != nil {
		panic(err)
	}

	fmt.Printf("Published image to: %s\n", ref)
}
