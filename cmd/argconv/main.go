package main

import "github.com/unmango/go/cli"

func main() {
	if err := Execute(); err != nil {
		cli.Fail(err)
	}
}
