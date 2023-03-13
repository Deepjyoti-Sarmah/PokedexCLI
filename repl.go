package main

import (
	"bufio"
	"fmt"
	"os"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf(" >")

	scanner.Scan()
	text := scanner.Text()

	fmt.Println("echoing:", text)

}