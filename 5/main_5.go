package main

import "fmt"

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}

	ok, m := ExistUnion(a, b)

	if ok {
			fmt.Printf("Значения пересечения двух слайсов %d\n", m)
	} else {
    fmt.Printf("Пересечений слайсов нет %d\n", m)
	} 
}

func ExistUnion(s1 []int, s2 []int) (bool, []int) {
	var exist bool = true

	if len(s1) == 0 || len(s2) == 0 {  
		exist = false                                                                      
    return exist, [] int{}                                                                                                 
  } 

	// мапа для быстрого доступа к эё элементам
  m := make(map[int]bool, 0)

	// Добавляем в неё элементы из первого слайса, и сравниваем с элементами из второго слайса
	for _, v1 := range s1 {
	  m[v1] = true
	}
 
  result := []int{}

	for _, v2 := range s2 {
		if m[v2] {
			result = append(result, v2)
		}
	}
	
	return exist, result
}
