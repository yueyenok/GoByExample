package main

import (
	"fmt"
)

func funcfor() {
	i := 0
	for i <= 6 {
		fmt.Println(i)
		i++
	}

	for j := 0; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("this is in a loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
