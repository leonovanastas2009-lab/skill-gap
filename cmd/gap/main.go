package main

import (
	"fmt"
	"os"
)

const version = "0.1"

func main() {
	if len(os.Args) < 2 {
		return
	}

	arg := os.Args[1]

	if arg == "hello" {
		fmt.Println("hello, brother — это gap v" + version)
		
	}else{
		fmt.Println("got:", arg)
	}

}
