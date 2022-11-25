package main

import (
	"fmt"
	"math/rand"
	"time"
)

func result(x int) bool {

	return x%2 == 0

}

func main() {
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(10)
	Sum := result(x)
	fmt.Println("Рандомное число", x)
	fmt.Println(Sum)

}
