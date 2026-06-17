package main

import (
	"fmt"
	"os"
	"strings"
)

const version = "0.1"

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	arg := os.Args[1]

	if arg == "hello" {
		fmt.Println("hello, brother — это gap v" + version)
		
	}else if arg == "echo" {
		if len(os.Args) < 3 {
			usage()
			os.Exit(1)
		}
		fmt.Println(strings.Join(os.Args[2:], " "))
	}else{
		usage()
		os.Exit(1)
	} 
}
func usage() {
	fmt.Fprintf(os.Stderr, "usage:\n")
	fmt.Fprintf(os.Stderr, " gap hello\n" )
	fmt.Fprintf(os.Stderr, " gap echo <text>\n")
}