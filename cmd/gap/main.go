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
		matchflags := flag.NewFlagSet("match", flag.ContinueOnError)
		id := matchflags.Int64("id", 0, "id match")

		if err := matchflags.Parse(os.Args[2:]); err != nil {

			os.Exit(1)
		}

		seen := map[string]bool{}
		matchflags.Visit(func(f *flag.Flag) { seen[f.Name] = true })

		if !seen["id"] {
			fmt.Fprintln(os.Stderr, "id flag is required")
			os.Exit(1)
		}
		if *id <= 0 {
			fmt.Fprintln(os.Stderr, "invalid id")
			os.Exit(1)
		}

		fmt.Fprintf(os.Stdout, "ID match = %d\n", *id)

	case "matches":
		matchesflags := flag.NewFlagSet("matches", flag.ContinueOnError)
		last := matchesflags.Int64("last", 10, "last matches")

		if err := matchesflags.Parse(os.Args[2:]); err != nil {
			os.Exit(1)
		}

		if *last <= 0 {
			fmt.Fprintln(os.Stderr, "limit is 0")
			os.Exit(1)
		}

		if *last > 50 {
			fmt.Fprintln(os.Stderr, "limit is 50")
			os.Exit(1)
		}
		fmt.Printf("matches для 12345678, last=%d (stub)\n", *last)

	case "player":
		playerflags := flag.NewFlagSet("player", flag.ContinueOnError)
		id := playerflags.Int64("id", 0, "id account")

		if err := playerflags.Parse(os.Args[2:]); err != nil {
			os.Exit(1)
		}

		if *id <= 0 {
			fmt.Fprintln(os.Stderr, "invalid id")
			os.Exit(1)

		}
		fmt.Fprintf(os.Stdout, "ID Игока = %d\n", *id)

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
	fmt.Fprintln(os.Stderr, "match --id=...")
	fmt.Fprintln(os.Stderr, "matches --last=...")
	fmt.Fprintln(os.Stderr, "player --id=...")

}
