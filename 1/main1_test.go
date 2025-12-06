package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"testing"
)

func TestConvertToStringsVar(t *testing.T) {
	result, err := convertVarsToString()
	if err != nil {
		t.Errorf("Ошибка при конвертации переменных: %v\n", err)
	}

	expectedStrNumDecimal := strconv.Itoa(numDecimal)
	expectedStrNumOctal := strconv.Itoa(numOctal)
	expectedStrNumHexadecimal := strconv.Itoa(numHexadecimal)
	expectedStrPi := strconv.FormatFloat(pi, 'f', -1, 64)
	expectedStrName := name
	expectedStrIsActive := strconv.FormatBool(isActive)
	expectedStrComplexNum := fmt.Sprintf("%v", complexNum)

	expected := expectedStrNumDecimal + expectedStrNumOctal + expectedStrNumHexadecimal +
		expectedStrPi + expectedStrName + expectedStrIsActive + expectedStrComplexNum

	if result == "" {
		t.Error("convertToStringVarsWithSalt() вернула пустую строку")
	}

	if result != expected {
		t.Errorf("convertToStringVarsWithSalt() вернула неожиданный результат.\nПолучено: %q\nОжидалось: %q", result, expected)
	}
}

func TestHashStringWithSalt(t *testing.T) {
	testString := "Hello World"
	testSalt := "go-2024"

	runes := []rune(testString)
	saltRunes := []rune(testSalt)

	mid := len(runes) / 2

	expectedRunes := make([]rune, 0, len(runes)+len(saltRunes))
	expectedRunes = append(expectedRunes, runes[:mid]...)
	expectedRunes = append(expectedRunes, saltRunes...)
	expectedRunes = append(expectedRunes, runes[mid:]...)

	expectedString := string(expectedRunes)
	expectedBytes := []byte(expectedString)

	expectedHash := sha256.Sum256(expectedBytes)

	resultHash, err := hashStringWithSalt(testString, testSalt)
	if err != nil {
		t.Errorf("Ошибка при хешировании: %v\n", err)
	}

	if len(resultHash) == 0 {
		t.Error("hashStringWithSalt() вернула пустой хеш")
	}

	if hex.EncodeToString(resultHash) != hex.EncodeToString(expectedHash[:]) {
		t.Errorf("hashStringWithSalt() вернула неправильный хеш.\nПолучено: %s\nОжидалось: %s",
			hex.EncodeToString(resultHash),
			hex.EncodeToString(expectedHash[:]))
	}
}
