package main
import "fmt"
func dan(a int){
    var i int
    i = 1
    fmt.Println(a)
    for i < 10 {
        fmt.Println(a, " * ", i, " = ", a * i )
        i = i + 1
    }
}

func main(){
    var j int
    j = 2
    for j < 10 {
        dan(j)
        j = j + 1
    }
}

