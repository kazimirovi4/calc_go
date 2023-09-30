package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var numbers = map[int]string{
	1:   "I",
	2:   "II",
	3:   "III",
	4:   "IV",
	5:   "V",
	6:   "VI",
	7:   "VII",
	8:   "VIII",
	9:   "IX",
	10:  "X",
	20:  "XX",
	30:  "XXX",
	40:  "XL",
	50:  "L",
	60:  "LX",
	70:  "LXX",
	80:  "LXXX",
	90:  "XC",
	100: "C",
}

func main() {
	fmt.Println("Введите операвцию: ")
	var num1, num2, operator string
	var num11, num22 int

	operators := [4]string{"+", "-", "*", "/"}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()

	str = strings.TrimSpace(str)

	num1_Ok := false

	for i := 0; i < len(str); i++ {
		if string(str[i]) != " " && !stringInSlice(string(str[i]), operators) && !num1_Ok {
			num1 += string(str[i])
		} else if (string(str[i]) == " " || stringInSlice(string(str[i]), operators)) && !num1_Ok {
			num1_Ok = true
			if string(str[i]) != " " {
				operator = string(str[i])
			}
		} else if stringInSlice(string(str[i]), operators) && num1_Ok && operator == "" {
			operator = string(str[i])
		} else if string(str[i]) != " " && num1_Ok {
			num2 += string(str[i])
		} else if string(str[i]) == " " && num1_Ok && num2 != "" {
			fmt.Println("Формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
			return
		}
	}

	if operator == "" {
		fmt.Println("Строка не является математической операцией.")
		return
	}

	rome1 := false
	rome2 := false

	num11, err := strconv.Atoi(num1)
	if err != nil {
		if toArabic(num1) > 0 {
			num11 = toArabic(num1)
			rome1 = true
		} else {
			fmt.Println("Разрешено использовать только целые числа от 1 до 10 включительно.")
			return
		}
	}

	num22, err = strconv.Atoi(num2)
	if err != nil {
		if toArabic(num2) > 0 {
			num22 = toArabic(num2)
			rome2 = true
		} else {
			fmt.Println("Разрешено использовать только целые числа от 1 до 10 включительно.")
			return
		}
	}

	if rome1 != rome2 {
		fmt.Println("Используются одновременно разные системы счисления.")
		return
	}

	if num11 < 1 || num11 > 10 || num22 < 1 || num22 > 10 {
		fmt.Println("Разрешено использовать только целые числа от 1 до 10 включительно.")
		return
	}

	if operator == "+" {
		if rome1 && rome2 {
			fmt.Println(toRome(num11 + num22))
		} else {
			fmt.Println(num11 + num22)
		}
	} else if operator == "-" {
		if rome1 && rome2 && num11-num22 < 1 {
			fmt.Println("В римской системе только положительные числа.")
		} else if rome1 && rome2 && num11-num22 > 0 {
			fmt.Println(toRome(num11 - num22))
		} else {
			fmt.Println(num11 - num22)
		}
	} else if operator == "/" {
		if rome1 && rome2 {
			if num11/num22 > 0 {
				fmt.Println(toRome(num11 / num22))
			} else {
				fmt.Println("ВВ римской системе только положительные числа.")
			}
		} else {
			fmt.Println(num11 / num22)
		}
	} else if operator == "*" {
		if rome1 && rome2 {
			fmt.Println(toRome(num11 * num22))
		} else {
			fmt.Println(num11 * num22)
		}
	}
}

func toRome(arab int) string {
	roman := ""
	if val, ok := numbers[arab]; ok {
		roman = val
	} else {
		ten := arab / 10 * 10
		if val, ok := numbers[ten]; ok {
			roman += val
		}
		remains := arab % 10
		if remains > 0 {
			if val, ok := numbers[remains]; ok {
				roman += val
			}
		}
	}
	return roman
}

func toArabic(roman string) int {
	arab := 0
	for key, value := range numbers {
		if value == roman {
			arab = key
		}
	}

	return arab
}

func stringInSlice(a string, arr [4]string) bool {
	for _, b := range arr {
		if b == a {
			return true
		}
	}
	return false
}
