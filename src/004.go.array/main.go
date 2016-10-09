package main
import "fmt"
func main() {
    x := [5]float64{
        1,
        1,
        1,
        1,
        1,
    }
    var total float64 = 0
    for _, value := range x {
        total += value
        fmt.Println(total)
    }
    fmt.Println(total / float64(len(x)))
}
