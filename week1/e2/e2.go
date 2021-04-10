package main

import "fmt"

func main() {

	arry := [7]int{1, 2, 3, 4, 5, 6, 7}

	for i, v := range arry {
		fmt.Println(v, i)
	}
	//better practice to use _
	for _, v := range arry {
		fmt.Println(v)
	}

	i := 0
	for {
		if i == 5 {
			break
		}
		i++
	}
	fmt.Println(i)
}
