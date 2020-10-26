package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	emailPtr := flag.String("email", "", "Email to ping when command completes.")
	flag.Parse()

	if *emailPtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	command := strings.Join(flag.Args(), " ")
	fmt.Printf("Email %s\nCommand %s\n", *emailPtr, command)
}
