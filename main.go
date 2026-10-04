package main

import "os"

const version = "0.1.0"

func main() {
	os.Exit(run(os.Args[1:]))
}
