package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// function to extract number next to cap,low,up - (for part 7)
func extract(look string, into []string) (int, error) {
	snew := string('(') + look + string(',')
	var errg error
	for i, slice := range into {
		if strings.Compare(slice, snew) == 0 {
			len := len(into[i+1])
			num := into[i+1][:len-1]
			a, err := strconv.Atoi(num)
			errg = err
			if err == nil && a > 0 {
				return a, err
			}
		}
	}
	return -1, errg
}

// function to check for vowel and "h"
func vowelh(str string) bool {
	runestr := []rune(str)
	a := runestr[0]
	if a == 'a' || a == 'e' || a == 'i' || a == 'o' || a == 'u' || a == 'h' {
		return true
	}
	return false
}

// function to check for puncts
func ispunct(a rune) bool {
	if a == '.' || a == ',' || a == '!' || a == '?' || a == ':' || a == ';' {
		return true
	}
	return false
}

func ispuncts(a string) bool {
	if a == "." || a == "," || a == "!" || a == "?" || a == ":" || a == ";" {
		return true
	}
	return false
}

// function to remove tags
func removetags(s []string) string {
	str := ""
	for i, tag := range s {
		if tag == "(cap," || tag == "(low," || tag == "(up," {
			s[i] = ""
			s[i+1] = ""
		} else if tag != "(up)" && tag != "(hex)" && tag != "(bin)" && tag != "(cap)" && tag != "(low)" && tag != "" {
			if i == 0 {
				str = str + tag
			} else {
				str = str + " " + tag
			}
		}
	}
	return str
}

// checking even number
func even(number int) bool {
	return number%2 == 0
}

func main() {
	// 1. Reading the file - receiving argument
	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	input := string(file) // the input
	slice := strings.Fields(input)
	lens := len(slice) - 1
	for i, first := range slice {
		// 2. hex: converting hexadecimal to decimal version / removing hex
		if strings.Compare(first, "(hex)") == 0 {
			num, _ := strconv.ParseInt(slice[i-1], 16, 64)
			slice[i-1] = strconv.Itoa(int(num))
		}
		// 3. bin: converting binary to decimal version / removing bin
		if strings.Compare(first, "(bin)") == 0 {
			num, _ := strconv.ParseInt(slice[i-1], 2, 64)
			slice[i-1] = fmt.Sprint(num)
		}
		// 4. up: converting uppercase
		if strings.Compare(first, "(up)") == 0 {
			slice[i-1] = strings.ToUpper(slice[i-1])
		}
		// 5. low: converting lowercase
		if strings.Compare(first, "(low)") == 0 {
			slice[i-1] = strings.ToLower(slice[i-1])
		}
		// 6. cap: capitalized first letter
		if strings.Compare(first, "(cap)") == 0 {
			slice[i-1] = strings.Title(slice[i-1])
		}
		// 7. up,low,cap, (number): transfer (number) of words before up,low,cap
		// 7.a. for cap
		if strings.Compare(first, ("(cap,")) == 0 {
			number, err := extract("cap", slice)
			if err == nil && i >= number { // Checks if the number of string before "cap" is enough to convert
				n := number
				for j := 1; j <= n; j++ {
					slice[i-j] = strings.Title(slice[i-j])
				}
			}
		}
		// 7.b. for up
		if strings.Compare(first, ("(up,")) == 0 {
			number, err := extract("up", slice)
			if err == nil && i >= number { // Checks if the number of string before "up" is enough to convert
				n := number
				for j := 1; j <= n; j++ {
					slice[i-j] = strings.ToUpper(slice[i-j])
				}
			}
		}
		// 7.c. for low
		if strings.Compare(first, ("(low,")) == 0 {
			number, err := extract("low", slice)
			if err == nil && i >= number { // Checks if the number of string before "low" is enough to convert
				n := number
				for j := 1; j <= n; j++ {
					slice[i-j] = strings.ToLower(slice[i-j])
				}
			}
		}
		// 8. a - an :vowel and "h"
		if i != lens {
			if slice[i] == "a" && vowelh(slice[i+1]) == true {
				slice[i] = slice[i] + "n"
			} else if slice[i] == "A" && vowelh(slice[i+1]) == true {
				slice[i] = slice[i] + "n"
			}
		}
	}
	output := removetags(slice)

	// 9.a correction of the punctuations' location
	str := []rune(output)
	lent := len(str)
	nquotes := 1
	for j := 0; j < lent-3; j++ {
		if j > 0 && str[j-1] == rune(39) {
			nquotes++
		}
		if ispunct(str[j+1]) == true {
			if str[j] == ' ' {
				str[j] = str[j+1]
				str[j+1] = ' '
			}
		}
		// 9.b. correction of the quotes
		if even(nquotes) == false && str[j] == rune(39) && str[j+1] == ' ' {
			str = append(str[0:j+1], str[j+2:]...)
		}
		if even(nquotes) == true && str[j+1] == rune(39) && str[j] == ' ' {
			str = append(str[0:j], str[j+1:]...)
		}
	}
	output = string(str)
	outslice := strings.Fields(output)
	l := len(outslice) - 1
	for i := 0; i <= l; i++ {
		if i == l && (outslice[i] == "'" || ispuncts(outslice[i]) == true) {
			outslice[i-1] = outslice[i-1] + outslice[i]
			outslice[i] = ""
		}
	}
	output = strings.Join(outslice, " ")
	output = strings.Join(strings.Fields(strings.TrimSpace(output)), " ")
	output2 := []byte(output)
	data := os.Args[2]
	errend := os.WriteFile(data, output2, 0o777)
	if errend != nil {
		fmt.Println(errend)
	}
}
