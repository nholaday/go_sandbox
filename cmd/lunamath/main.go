package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var counter uint = 1
	var total uint = 1

	fmt.Println("Welcome to LunaMath!  Press return to generate the next number")
	for true {
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		fmt.Printf("counter: %v, total %v", counter, total)
		fmt.Println(" -- Press return again to generate the next number")
		counter++
		total += counter
	}
}