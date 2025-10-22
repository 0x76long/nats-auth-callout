package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/nats-io/jwt/v2"
	"github.com/nats-io/nats.go"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	flag.Parse()

	// Model the user encoded in the users file.
	type User struct {
		Pass        string
		Account     string
		Permissions jwt.Permissions
	}

	// Open the NATS connection passing the auth account creds file.
	nc, err := nats.Connect("nats://test:password@localhost")
	if err != nil {
		return err
	}
	defer nc.Drain()

	fmt.Println("test")

	return nil
}
