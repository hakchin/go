// factorial calculation
package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)
	x := new(big.Int)
	var lv_input int64
	lv_input = int64(input)
	x.MulRange(1, lv_input)
	fmt.Println(x)
}
