
package main

import (
	"fmt"
	"bufio"
	"os"
	s "strings"
)

func handleInput(input string, details string) {
	if s.Contains(input, "testing") == true {
		fmt.Println("/c game.print('Got input: ", input, "')")
	} else {
		fmt.Println("Invalid :  ", input)
	}
}

func main() {
	fmt.Println("/c game.print('Clusterio Example Plugin enabled!')")
	waitForInput()
}

func waitForInput() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		//fmt.Println(s)
		handleInput(input, "")
	}
	if scanner.Err() != nil {

	}
}
