package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getnumber(x, y int) (int, int) {
	x = 2*x + 10
	y = -3*y - 5
	return x, y
}
func getRandom(a, b int) (int, int) {
	return rand.Intn(a), rand.Intn(b)
}

func main() {

	rand.Seed(time.Now().UnixNano())

	x1, y1 := getnumber(getRandom(15, 25))
	x2, y2 := getnumber(getRandom(35, 42))
	x3, y3 := getnumber(getRandom(10, 21))
	fmt.Println(x1, y1)
	fmt.Println(x2, y2)
	fmt.Println(x3, y3)
}
