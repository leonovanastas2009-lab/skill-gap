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
	case "match":
		matchflags := flag.NewFlagSet("match", flag.ExitOnError)
		id := matchflags.Int("id", 0, "id match")

		if err := matchflags.Parse(os.Args[2:]); err != nil {
			os.Exit(1)
		}

		if *id <= 0 {
			fmt.Fprint(os.Stderr, "invalid id")
			os.Exit(1)

		} else {
			fmt.Fprintf(os.Stdout, "ID match = %d\n", *id)
		}

	case "matches":
		matchesflags := flag.NewFlagSet("matches", flag.ExitOnError)
		last := matchesflags.Int("last", 10, "last matches")

		if err := matchesflags.Parse(os.Args[2:]); err != nil {
			os.Exit(1)
		}

		if *last > 50 {
			fmt.Fprint(os.Stderr, "Limit is 50")
			os.Exit(1)
		} else {
			fmt.Fprintf(os.Stdout, "matches для 12345678, last= %d (stub)", *last)
		}

	case "player":
		playerflags := flag.NewFlagSet("player", flag.ExitOnError)
		id := playerflags.Int("id", 0, "id accaunt")

		if err := playerflags.Parse(os.Args[2:]); err != nil {
			os.Exit(1)
		}

		if *id <= 0 {
			fmt.Fprint(os.Stderr, "invalid id")
			os.Exit(1)

		} else {
			fmt.Fprintf(os.Stdout, "ID Игока = %d\n", *id)
		}

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
	fmt.Fprintln(os.Stderr, "match --id=...")
	fmt.Fprintln(os.Stderr, "matches --last=...")
	fmt.Fprintln(os.Stderr, "player --id=...")
	fmt.Fprintln(os.Stderr, "usage:")
	fmt.Fprintln(os.Stderr, "gap hello")
	fmt.Fprintln(os.Stderr, "gap echo ...")
}
