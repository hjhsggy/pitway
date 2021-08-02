package main

import (
	"fmt"
	"os"
)

func main() {

	ch := make(chan os.Signal, 1)
	// for i := 0; i < 10; i++ {
	// 	go func() {
	// 		fmt.Println(i)
	// 	}()
	// }

	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}

	select {
	case ch <- os.Interrupt:
		break
	}

}
