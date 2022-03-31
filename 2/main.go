package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func UnpackString(s string) (string, error) {
	if s == "" {
		return "", nil
	}

	prev := ""
	runes := []rune(s)
	length := len(runes)
	result := ""

	flag := false
	num := 0

	if unicode.IsDigit(runes[0]) {
		err := fmt.Errorf("Invalid string %s ", s)
		return "", err
	}

	for i := 0; i < length; i++ {
		if runes[i] == '\\' && !flag {
			flag = true
			continue
		}

		switch {
		case flag && prev == "":
			prev = string(runes[i])
			flag = false
		case !flag && unicode.IsDigit(runes[i]):
			digit, err := strconv.Atoi(string(runes[i]))
			if err != nil {
				return "", err
			}
			num = num*10 + digit
			if i+1 == length || !unicode.IsDigit(runes[i+1]) {
				result = addToStr(result, prev, num)
				num = 0
				prev = ""
			}
		case !flag && !unicode.IsLetter(runes[i]):
			err := fmt.Errorf("Invalid string %s ", s)
			return "", err
		default:
			result = addToStr(result, prev, 1)
			prev = string(runes[i])
			flag = false
		}
	}

	if flag {
		err := fmt.Errorf("Invalid string %s ", s)
		return "", err
	}

	result = addToStr(result, prev, 1)

	return result, nil
}

func addToStr(word string, symbol string, count int) string {
	for i := 0; i < count; i++ {
		word += symbol
	}
	return word
}

func main() {
	fmt.Println(UnpackString("g13j4"))
	fmt.Println(UnpackString("a4bc2d5e"))
	fmt.Println(UnpackString("abcd"))
	fmt.Println(UnpackString("45"))
	fmt.Println(UnpackString(""))
	fmt.Println(UnpackString(`qwe\4\5`))
	fmt.Println(UnpackString(`qwe\45`))
	fmt.Println(UnpackString(`qwe\\5`))
	fmt.Println(UnpackString(`5\4`))
	fmt.Println(UnpackString(`\\_`))
	fmt.Println(UnpackString(`\\\`))
}
