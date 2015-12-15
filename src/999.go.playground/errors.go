package main
import "fmt"
type Foo string
func (Foo) Error() string {
    return "error"
}
func (Foo) String() string {
    return "string"
}

func main() {
    f := Foo("foo")
    fmt.Printf("%s",f)
}
