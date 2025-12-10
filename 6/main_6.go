package main

import (
	"fmt"
	"math/rand"
)

func main() {
	ch := make(chan(int))
  go GenRandomIntWithChan(ch)
	v := <-ch
	fmt.Println("Поучили из канала число")
	fmt.Println(v)
}

func GenRandomIntWithChan(ch chan int) {
	num := rand.Intn(1000)
	ch <- num
}
