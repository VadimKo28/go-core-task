package main

import (
	"fmt"
	"math/rand"
)

func main() {
	size := 10
	s := make([]int, size)
	randomSliceNumbers := genenerateRandomNumbersSlice(size, s)
	fmt.Printf("Сгенерили слайс случайных чисел: %d\n", randomSliceNumbers)

  evenNumbers := sliceExample(s)
	fmt.Printf("Из исходного слайса получили новый слайс чётных чисел: %d\n", evenNumbers)

	fmt.Printf("элементы слайса: % d\n", s)
	newSlice := addElements(s, 999)
	fmt.Printf("Добавили новый элемент в слайс: % d\n", newSlice)

	fmt.Printf("Передаём в функцию copy слайс %d\n", s)
	fmt.Printf("Его адрес %p\n", s)

	copy := copySlice(s)
	fmt.Printf("Получаем копию слайса %d\n", copy)
	fmt.Printf("Её адрес %p\n", copy)

	fmt.Printf("Передаём в функцию удаления слайс %d\n", s)
	index := 4
	newSlise := removeElement(s, index)
  fmt.Printf("Получаем новый слайс без элемента с индексом %d\n", index)
	fmt.Println(newSlise)
}

func addElements(s []int, element int) []int {
  s = append(s, element)
	return s
}

func copySlice(s []int) []int {
  dst := make([]int, len(s))

  copy(dst, s)

	return dst
}

func genenerateRandomNumbersSlice(size int, s []int) []int {
	for i := range size {
		s[i] = rand.Intn(100)
	}

	return s
}

func removeElement(s []int, index int) []int {
	if index < 0 || index >= len(s) {
		return s
	}

	return append(s[:index], s[index+1:]...)
}

func sliceExample(s []int) []int {
	newSlice := make([]int, 0)

	for _, value := range s  {
    if value % 2 == 0 {
      newSlice = append(newSlice, value)
		}
	}

  return newSlice
}
