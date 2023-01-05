package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter data:")
	reader := bufio.NewReader(os.Stdin)
	taskText, _ := reader.ReadString('\n')
	taskText = strings.ReplaceAll(strings.TrimSpace(taskText), " ", "")
	calculate(taskText)
}

func calculate(taskText string) {
	var charAction string
	var err error

	if strings.Contains(taskText, "+") {
		charAction = "+"
	} else if strings.Contains(taskText, "-") {
		charAction = "-"
	} else if strings.Contains(taskText, "/") {
		charAction = "/"
	} else if strings.Contains(taskText, "*") {
		charAction = "*"
	} else {
		err = errors.New("Invalid data format")
	}

	if err != nil {
		causeError(err)
		return
	}

	arrNums := strings.Split(taskText, charAction)

	if len(arrNums) != 2 {
		causeError(errors.New("Invalid data format"))
		return
	}

	romanNums := false
	var res, num1, num2 int
	num1, err = strconv.Atoi(arrNums[0])

	if err != nil {
		num1, err = convertRomanToInt(arrNums[0])

		if err != nil {
			causeError(err)
			return
		} else if num1 < 1 || num1 > 10 {
			causeError(errors.New("Invalid number values"))
			return
		} else {
			romanNums = true
		}
	}

	if romanNums {
		num2, err = convertRomanToInt(arrNums[1])
	} else {
		num2, err = strconv.Atoi(arrNums[1])
	}

	if err != nil {
		causeError(err)
		return
	} else if num1 < 1 || num1 > 10 || num2 < 1 || num2 > 10 {
		causeError(errors.New("Invalid number values"))
		return
	}

	switch charAction {
	case "+":
		res = num1 + num2
	case "-":
		res = num1 - num2
	case "/":
		res = num1 / num2
	case "*":
		res = num1 * num2
	}

	if romanNums {
		if res < 1 {
			causeError(errors.New("Invalid number values"))
			return
		} else {
			fmt.Println(convertIntToRoman(res))
		}
	} else {
		fmt.Println(res)
	}
}

func causeError(err error) {
	fmt.Println(err)
}

func convertRomanToInt(numRomanStr string) (int, error) {
	arrRune := []rune(numRomanStr)
	romanNums := map[rune]int{'I': 1, 'V': 5, 'X': 10}
	lenStr := len(arrRune) - 1
	numRes := romanNums[arrRune[lenStr]]

	if numRes == 0 {
		return 0, errors.New("Invalid data format")
	}

	for lenStr > 0 {
		if arrRune[lenStr] == rune('I') {
			if arrRune[lenStr-1] == rune('I') || arrRune[lenStr-1] == rune('V') || arrRune[lenStr-1] == rune('X') {
				numRes += romanNums[arrRune[lenStr-1]]
			} else {
				return 0, errors.New("Invalid data format")
			}
		} else if arrRune[lenStr] == rune('V') {
			if arrRune[lenStr-1] == rune('V') || arrRune[lenStr-1] == rune('X') {
				numRes += romanNums[arrRune[lenStr-1]]
			} else if arrRune[lenStr-1] == rune('I') {
				numRes -= romanNums[arrRune[lenStr-1]]
			} else {
				return 0, errors.New("Invalid data format")
			}
		} else if arrRune[lenStr] == rune('X') {
			if arrRune[lenStr-1] == rune('X') {
				numRes += romanNums[arrRune[lenStr-1]]
			} else if arrRune[lenStr-1] == rune('I') || arrRune[lenStr-1] == rune('V') {
				numRes -= romanNums[arrRune[lenStr-1]]
			} else {
				return 0, errors.New("Invalid data format")
			}
		} else {
			return 0, errors.New("Invalid data format")
		}

		lenStr--
	}

	return numRes, nil
}

func convertIntToRoman(num int) string {
	romanNums := map[int]rune{1: 'I', 5: 'V', 10: 'X', 50: 'L', 100: 'C'}
	var numRuneRoman []rune

	number10 := num / 10

	if number10 == 10 {
		numRuneRoman = append(numRuneRoman, romanNums[100])
	} else if number10 == 5 {
		numRuneRoman = append(numRuneRoman, romanNums[50])
	} else if number10 > 5 {
		numRuneRoman = append(numRuneRoman, romanNums[50])
		for i := 0; i < number10-5; i++ {
			numRuneRoman = append(numRuneRoman, romanNums[10])
		}
	} else if number10 != 0 {
		for i := 0; i < number10; i++ {
			numRuneRoman = append(numRuneRoman, romanNums[10])
		}
	}

	number1 := num % 10

	if number1 == 5 {
		numRuneRoman = append(numRuneRoman, romanNums[5])
	} else if number1 == 4 {
		numRuneRoman = append(numRuneRoman, romanNums[1], romanNums[5])
	} else if number1 == 9 {
		numRuneRoman = append(numRuneRoman, romanNums[1], romanNums[10])
	} else if number1 > 0 && number1 < 4 {
		for i := 0; i < number1; i++ {
			numRuneRoman = append(numRuneRoman, romanNums[1])
		}
	} else if number1 > 5 && number1 < 9 {
		numRuneRoman = append(numRuneRoman, romanNums[5])
		for i := 0; i < number1-5; i++ {
			numRuneRoman = append(numRuneRoman, romanNums[1])
		}
	}

	return string(numRuneRoman)
}
