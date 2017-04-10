// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	// "strings"
)

func main() {
	/*
		    s, sep := "", ""
			for _, arg := range os.Args[1:] {
				s += sep + arg
				sep = " "
			}
			fmt.Println(s)
	*/
	// fmt.Println(strings.Join(os.Args[1:], " ")) // arg1 arg2 ...
	// fmt.Println(strings.Join(os.Args, " ")) // exe arg1 arg2 ...

	for idx, arg := range os.Args[1:] {
		fmt.Printf("%d : %v\n", idx, arg)
	}
}
