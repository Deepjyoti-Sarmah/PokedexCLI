package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print(" >")

	scanner.Scan()
	test := scanner.Text()

	fmt.Println("echoing: ", test)
}
