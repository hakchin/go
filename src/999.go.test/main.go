package main

import "fmt"

func main() {
	numberMap := map[string]int{}
	numberMap["zero"] = 0
	numberMap["one"] = 1
	numberMap["two"] = 2
	fmt.Println(numberMap["zero"])
	fmt.Println(numberMap["one"])
	fmt.Println(numberMap["two"])

	if v, ok := numberMap["three"]; ok {
		fmt.Println("three' is in numberMap.value:", v)
	} else {
		fmt.Println("three' is not in numberMap")
		fmt.Println(ok)
	}
	if v, ok := numberMap["two"]; ok {
		fmt.Println("two' is in numberMap.value:", v)
		fmt.Println(ok)
	} else {
		fmt.Println("two' is not in numberMap")
	}
}
