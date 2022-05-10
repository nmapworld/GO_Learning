package main

import (
	"fmt"
	"os"
)

// args
func main() {
	fmt.Println("hello wolrd")
	if len(os.Args) > 0 {
		for index, arg := range os.Args {
			fmt.Println(index, arg)
		}
	}

}
