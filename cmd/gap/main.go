package main

import (
	"flag"
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
		helloflags := flag.NewFlagSet("hello", flag.ExitOnError)
		name := helloflags.String("name", "brother", "имя")

		if err := helloflags.Parse(os.Args[2:]); err != nil {
			os.Exit(1)
		}

		fmt.Fprintf(os.Stdout, "Привет, %s это gap v%s\n", *name, version)

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
