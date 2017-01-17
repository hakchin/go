// usage : go run main.go 36.5
// Cf converts its numeric argument to Celsius and Fahrenheit.
package main

import (
	"com/ppoppo/tempconv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
	}

	i, err := strconv.Atoi("-42")
	s := strconv.Itoa(-42)
	fmt.Println(i)
	fmt.Println(s)
	fmt.Println(err)
}
