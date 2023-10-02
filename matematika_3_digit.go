package main

import "fmt"

func main() {
	var x, v1, v2, v3, hb int32
	fmt.Scanln(&x)
	v1 = x / 100
	hb = x % 100
	v2 = hb / 10
	v3 = hb % 10

	fmt.Println(v1, v2, v3)
}
