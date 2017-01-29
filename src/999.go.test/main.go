package main

import "fmt"

func main() {
	myFunc("sjfskladjfs", 898, 8989)
}
func myFunc(s string, integers ...int) {
	fmt.Println(s)
	for i := 0; i < len(integers); i++ {
		fmt.Println(integers[i])
	}
}
