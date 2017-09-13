package main
import "C"

//export Hello
func Hello() string {
    return "Hello"
}

//export Test
func Test(){
    println("test");
}

func main() {
}