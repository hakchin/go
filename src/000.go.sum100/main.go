package main
import "fmt"
func main(){
    var i int
    var sum int
    i = 1
    sum = 0
    for i < 101 {
        sum = sum + i
        i = i + 1
    }
    fmt.Println("sum = ", sum)
}
