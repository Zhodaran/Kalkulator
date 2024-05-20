package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var roman = map[string]uint{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
}

var uintroman = map[uint]string{
	1:  "I",
	2:  "II",
	3:  "III",
	4:  "IV",
	5:  "V",
	6:  "VI",
	7:  "VII",
	8:  "VIII",
	9:  "IX",
	10: "X",
}

func sumRomaB(a, b string) string {
	valA, _ := roman[a]
	valB, _ := roman[b]
	sum := valA + valB
	if sum >= 30 && sum < 39 {
		return "X" + "X" + "X" + uintroman[sum-30]
	} else if sum >= 20 && sum < 29 {
		return "X" + "X" + uintroman[sum-20]
	} else if sum >= 10 && sum < 19 {
		return "X" + uintroman[sum-10]
	}
	result, _ := uintroman[sum]
	return result
}

func destRomaB(a, b string) string {
	valA, _ := roman[a]
	valB, _ := roman[b]
	if valA < valB {
		panic("В римской системе нет отрицательных числе")
	}
	sum := valA - valB
	if sum == 0 {
		panic("В римской системе нет 0")
	}

	if sum >= 30 && sum < 39 {
		return "X" + "X" + "X" + uintroman[sum-30]
	} else if sum >= 20 && sum < 29 {
		return "X" + "X" + uintroman[sum-20]
	} else if sum >= 10 && sum < 19 {
		return "X" + uintroman[sum-10]
	}

	result, _ := uintroman[sum]
	return result
}

func multRomaB(a, b string) string {
	valA, _ := roman[a]
	valB, _ := roman[b]
	sum := valA * valB
	if sum == 100 {
		return "C"
	} else if sum >= 90 && sum <= 99 {
		return "XC" + uintroman[sum-90]
	} else if sum >= 80 && sum <= 89 {
		return "LXXX" + uintroman[sum-80]
	} else if sum >= 70 && sum <= 79 {
		return "LXX" + uintroman[sum-70]
	} else if sum >= 60 && sum <= 69 {
		return "LX" + uintroman[sum-60]
	} else if sum >= 50 && sum <= 59 {
		return "L" + uintroman[sum-50]
	} else if sum >= 40 && sum <= 49 {
		return "XL" + uintroman[sum-40]
	} else if sum >= 30 && sum < 39 {
		return "X" + "X" + "X" + uintroman[sum-30]
	} else if sum >= 20 && sum < 29 {
		return "X" + "X" + uintroman[sum-20]
	} else if sum >= 10 && sum < 19 {
		return "X" + uintroman[sum-10]
	}
	result, _ := uintroman[sum]
	return result
}

func fractRomaB(a, b string) string {
	valA, _ := roman[a]
	valB, _ := roman[b]
	sum := valA / valB

	if sum == 0 {
		panic("В римской системе нет 0")
	}

	if sum >= 30 && sum < 39 {
		return "X" + "X" + "X" + uintroman[sum-30]
	} else if sum >= 20 && sum < 29 {
		return "X" + "X" + uintroman[sum-20]
	} else if sum >= 10 && sum < 19 {
		return "X" + uintroman[sum-10]
	}
	result, _ := uintroman[sum]
	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка")
		return
	}
	input = strings.TrimSpace(input)
	parts := strings.Fields(input)
	if len(parts) < 3 {
		panic("Строка не является математической операцией")
	}
	if len(parts) > 3 {
		panic("Ошибка: калькуляцию необходимо производить с двумя операндами")

	}

	num1 := parts[0]
	znak := parts[1]
	num2 := parts[2]
	/*input = strings.TrimSpace(input) // чистка пробелов
	parts := strings.Fields(input)   // разбивка на кол-во

	if len(parts) != 3 {
		fmt.Println("Вводимых зачений не должно быть больше двух")
		return
	}

	num1 := parts[0]
	znak := parts[1]
	num2 := parts[2]
	*/
	if (num1 == "I" || num1 == "II" || num1 == "III" || num1 == "IV" || num1 == "V" || num1 == "VI" || num1 == "VII" || num1 == "VIII" || num1 == "IX" || num1 == "X") && (num2 == "I" || num2 == "II" || num2 == "III" || num2 == "IV" || num2 == "V" || num2 == "VI" || num2 == "VII" || num2 == "VIII" || num2 == "IX" || num2 == "X") {
		if znak == "+" {
			fmt.Println(sumRomaB(num1, num2))
			return
		} else if znak == "-" {
			fmt.Println(destRomaB(num1, num2))
			return
		} else if znak == "*" {
			fmt.Println(multRomaB(num1, num2))
			return
		} else if znak == "/" {
			fmt.Println(fractRomaB(num1, num2))
			return
		}
	}
	number1, _ := strconv.Atoi(num1)
	number2, _ := strconv.Atoi(num2)
	if number1 > 10 || number2 > 10 {
		panic("Ошибка: числа не могут быть больше 10")
	}
	if (number1 >= 1 && number1 <= 10) && (number2 >= 1 && number2 <= 10) {

		if znak == "+" {
			fmt.Println(number1 + number2)
			return
		} else if znak == "-" {
			fmt.Println(number1 - number2)
			return
		} else if znak == "*" {
			fmt.Println(number1 * number2)
			return
		} else if znak == "/" {
			fmt.Println(number1 / number2)
			return
		}

	} else {
		panic("Римские и арабские цифры нельзя складывать")
	}

}
