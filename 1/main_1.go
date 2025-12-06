package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
)

var numDecimal int = 42
var numOctal int = 052
var numHexadecimal int = 0x2A
var pi float64 = 3.14
var name string = "Golang"
var isActive bool = true
var complexNum complex64 = 1 + 2i
var salt string = "go-2024"

func main() {
	printTypesVars()

	s, err := convertVarsToString()
	if err != nil {
		fmt.Printf("Ошибка при конвертации переменных: %v\n", err)
		return
	}

	hashed, err := hashStringWithSalt(s, salt)
	if err != nil {
		fmt.Printf("Ошибка при хешировании: %v\n", err)
		return
	}
	
	fmt.Printf("Хеш SHA256: %s\n", hashed)
}

func convertVarsToString() (string, error) {
	strNumDecimal := strconv.Itoa(numDecimal)
	strNumOctal := strconv.Itoa(numOctal)
	strNumHexadecimal := strconv.Itoa(numHexadecimal)
	strPi := strconv.FormatFloat(pi, 'f', -1, 64)
	strName := name
	strIsActive := strconv.FormatBool(isActive)
	strComplexNum := fmt.Sprintf("%v", complexNum)

	allStrings := strNumDecimal + strNumOctal + strNumHexadecimal + strPi + strName + strIsActive + strComplexNum

	return allStrings, nil
}

func hashStringWithSalt(str string, salt string) (string, error) {
	runes := []rune(str)
	saltRunes := []rune(salt)

	mid := len(runes) / 2

	// Собираем новый срез рун, соль помещаем в середине
	resultRunes := make([]rune, 0, len(runes)+len(saltRunes))
	resultRunes = append(resultRunes, runes[:mid]...)
	resultRunes = append(resultRunes, saltRunes...)
	resultRunes = append(resultRunes, runes[mid:]...)

	resultString := string(resultRunes)
	resultBytes := []byte(resultString)

	hasher := sha256.New()
	hasher.Write(resultBytes)

	return hex.EncodeToString(hasher.Sum(nil)), nil
}

func printTypesVars() {
	fmt.Printf("numDecimal type - %T\n", numDecimal)
	fmt.Printf("numOctal type - %T\n", numOctal)
	fmt.Printf("numHexadecimal type - %T\n", numHexadecimal)
	fmt.Printf("pi type - %T\n", pi)
	fmt.Printf("name type - %T\n", name)
	fmt.Printf("isActive type - %T\n", isActive)
	fmt.Printf("complexNum type - %T\n", complexNum)
}
