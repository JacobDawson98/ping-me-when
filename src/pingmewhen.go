package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"bytes"
	"log"
)

func main() {
	emailPtr := flag.String("email", "", "Email to ping when command completes.")
	flag.Parse()

	if *emailPtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Execute command
	givenCmd := strings.Join(flag.Args(), " ")
	var out bytes.Buffer
	cmd := exec.Command(givenCmd)
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Email %s\nCommand %s\n", *emailPtr, givenCmd)
	fmt.Printf(out.String())
}
