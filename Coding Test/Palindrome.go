package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("input kata :")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		line := scanner.Text()
		for i := 0; i < len(line)/2; i++ {
			if line[i] != line[len(line)-1-i] {
				fmt.Println("bukan kata palindrome")
			} else {
				fmt.Println("kata palindrome")
			}
		}

	}

}
