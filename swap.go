package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}
func main() {
	a, b := swap("Hello", "RamReddy")
	fmt.Println(a, b)
}
