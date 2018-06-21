package main;

import "fmt"

func main() {
    var result string
    src := "stressed"
    for i, _ := range src  {
        result += string(src[len(src) - i - 1])
    }
    fmt.Println(result)
}
