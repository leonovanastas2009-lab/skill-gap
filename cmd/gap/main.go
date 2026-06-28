package main

import (
	"fmt"
	"os"
)

const version = "0.1"

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "hello":
		fmt.Println("hello, brother это gap v" + version)
	case "echo":
		if len(os.Args) < 3 {
			usage()
			os.Exit(1)
		}

		fmt.Println(os.Args[2:])
	default:
		usage()
		os.Exit(1)
	}

}

func usage() {
	fmt.Fprintln(os.Stderr, "usage:")
	fmt.Fprintln(os.Stderr, "gap hello")
	fmt.Fprintln(os.Stderr, "gap echo ...")
}
