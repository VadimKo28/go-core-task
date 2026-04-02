package main

import "fmt"

func main() {
	slice1 := []string{"apple", "banana", "1", "2"}
	slice2 := []string{"apple", "orange", "1"}

	result := Difference(slice1, slice2)

	fmt.Printf("Эти элементы %s есть только в первом слайсе\n", result)
}


func Difference(slice1, slice2 []string) []string {
	set2 := make(map[string]bool)
	for _, v := range slice2 {
		set2[v] = true
	}

	result := make([]string, 0)
	for _, v := range slice1 {
		if !set2[v] {
			result = append(result, v)
		}
	}

	return result
}
