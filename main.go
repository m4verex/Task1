package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	text = strings.TrimSpace(text)

	textSplit := strings.Split(text, " ")

	if len(textSplit) > 3 {
		log.Panicln("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	}
	 if len(textSplit) <3 {
		log.Panicln("Выдача паники, так как строка не является математической операцией.")
	 }
	a, isRomanA := isDigit(textSplit[0])
	b, isRomanB := isDigit(textSplit[2])

	if isRomanA != isRomanB {
		log.Panicln("Выдача паники, так как используются одновременно разные системы счисления.")
	}

	switch textSplit[1] {
	case "+":
		if isRomanA {
			fmt.Println(intToRoman(a + b))
		} else {
			fmt.Println(a + b)
		}
	case "-":
		if isRomanA {
			fmt.Println(intToRoman(a - b))
		} else {
			fmt.Println(a - b)
		}
	case "*":
		if isRomanA {
			fmt.Println(intToRoman(a * b))
		} else {
			fmt.Println(a * b)
		}
	case "/":
		if isRomanA {
			fmt.Println(intToRoman(a / b))
		} else {
			fmt.Println(a / b)
		}
	}

}

func isDigit(str string) (res int, isRoman bool) {
	res, err := strconv.Atoi(str)
	if err != nil {
		res, err = romanToInt(str)
		isRoman = true
		if err != nil {
			log.Panicln(err)
		}
	}
	return res, isRoman
}

func romanToInt(s string) (result int, err error) {
	rMap := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	for k := range s {
		if k < len(s)-1 && rMap[s[k:k+1]] < rMap[s[k+1:k+2]] {
			result -= rMap[s[k:k+1]]
		} else {
			result += rMap[s[k:k+1]]
		}
	}
	if result == 0 {
		err = errors.New("Выдача паники, так как число не является римской цифрой")
		return result, err
	} else {
		return result, err
	}

}

func intToRoman(num int) string {
	if num < 0 {
		log.Panicln("Выдача паники, так как в римской системе нет отрицательных чисел.")
	}
	r := [][]string{
		{"", "M", "MM", "MMM"},
		{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"},
		{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"},
		{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}}
	n := []int{1000, 100, 10, 1}
	result := ""
	for k, v := range n {
		result += r[k][num/v]
		num = num % v
	}
	return result
}
