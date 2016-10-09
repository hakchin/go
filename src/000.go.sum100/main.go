package main
import "fmt"
func main(){
    var i int
    var sum int
    sum = 0
    for i = 1; i < 101 ; i++ {
        sum += i
    }
    fmt.Println("sum = ", sum)
}
