package main

import "fmt"

func main() {
	src := []rune("パタトクカシーー")
	var result string

	for i, _ := range src {
		if i%2 == 0 {
			result += string(src[i])
		}
	}

	fmt.Println(result)
}
