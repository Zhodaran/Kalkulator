package main

import "fmt"
import "strconv"
import "strings"
import "bufio"
import "os"

func splitStrings(inputStr string) (string, string) {
	delimiters := []string{" + ", " - ", " * ", " / "}
	for _, delimiter := range delimiters {
		if parts := strings.SplitN(inputStr, delimiter, 2); len(parts) == 2 {
			return strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
		}

	}
	return "", ""
}
func main() {
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	parts := strings.Fields(input)
	var t = len(parts)
	var znak string
	for i := 0; i < t; i++ {
		if parts[i] == "+" || parts[i] == "-" || parts[i] == "/" || parts[i] == "*" {
			znak = parts[i]
		}
	}
	str1, str2 := splitStrings(input)

	num2, _ := strconv.Atoi(str2)
	num1, _ := strconv.Atoi(str1)
	if num1 >= 1 && num1 <= 10 {
		panic("Первой строкой не может быть число")
	}

	if num2 > 10 {
		panic("Операнд больше 10")
	} else if strings.HasPrefix(str2, `"`) && strings.HasSuffix(str2, `"`) {
		str2 = str2[1 : len(str2)-0]

	} else if num2 < 10 {

	} else {
		panic("Строка без ковычек")
	}

	if strings.HasPrefix(str1, `"`) && strings.HasSuffix(str1, `"`) {
		str1 = str1[:len(str1)-1]
	} else if num1 > 10 {
		panic("Число больше 10")
	} else {
		fmt.Println(str1)
		panic("Строка без ковычек")
	}
	switch znak {
	case "*":
		if num2 >= 1 && num2 <= 10 {
			str := str1
			str = str[1 : len(str)-1]
			for i := 0; i < num2-1; i++ {
				str1 += str
			}
			fmt.Println(str1, `"`)
			return
		}

		fmt.Println(str1 + str2)
		return

	case "+":
		str1 += str2
		fmt.Println(str1)
		return
	case "/":
		strD := strings.Split(str1, "")
		var ss int
		ss = len(strD) / num2
		f := strD[0 : ss+1]
		resultD := strings.Join(f, "")
		fmt.Println(resultD, `"`)

		return

	case "-":
		/*   strF1 := strings.Split(str1, "")
		     strF2 := strings.Split(str2, "")
		     fmt.Println(strF1, strF2)
		     strFR := append(strF1, strF2...)
		     fmt.Println(strFR) */
		if strings.HasSuffix(str2, `"`) {
			str2 = str2[:len(str2)-1]
		}

		if strings.HasPrefix(str1, `"`) {
			str1 = str1[1:len(str1)]
		}

		result := strings.ReplaceAll(str1, str2, "")
		result += `"`
		result = `"` + result
		fmt.Println(result)
		return

	default:
		fmt.Println(str1)
	}
}
