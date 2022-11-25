package main

import (
	"fmt"
	"math/rand"
	"time"
)

func miltiply(x *int) {

	*x = *x * 10
	//Первое действие мы умножили и получили", x

}

func plus(x *int) {

	*x = *x + 10
	//"Второе действие прибавили и получили ", x

}

func main() {

	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(10)

	fmt.Println("Рандомное число равно", x)

	miltiply(&x)
	fmt.Println("Первое действие мы умножили и получили число", x)
	plus(&x)
	fmt.Println("Второе действие мы прибавили к первому и получили число", x)

}
