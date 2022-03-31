package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Flag struct {
	columnNumber  int
	isNumericSort bool
	isReverse     bool
	isOnlyUnique  bool
}

func main() {
	columnNumber := flag.Int("k", 0, "указание колонки для сортировки")
	isNumericSort := flag.Bool("n", false, "сортировать по числовому значению")
	isReverse := flag.Bool("r", false, "сортировать в обратном порядке")
	isOnlyUnique := flag.Bool("u", false, "не выводить повторяющиеся строки")

	flag.Parse()

	var input string
	bytes, err := os.ReadFile(flag.Arg(0))
	if err != nil {
		fmt.Println("error opening file: err:", err)
		os.Exit(1)
	}

	input = string(bytes)

	parameters := Flag{
		columnNumber:  *columnNumber,
		isNumericSort: *isNumericSort,
		isReverse:     *isReverse,
		isOnlyUnique:  *isOnlyUnique,
	}

	output, err := sortLines(input, parameters)
	if err != nil {
		fmt.Println(err)
	}
	for _, line := range output {
		fmt.Println(line)
	}

}

func sortLines(input string, parameters Flag) (result []string, err error) {

	listOfStrings := strings.Split(input, "\n")

	var sortedLines []string
	if parameters.isOnlyUnique {
		set := make(map[string]bool)
		for _, value := range listOfStrings {
			set[value] = true
		}

		for value := range set {
			sortedLines = append(sortedLines, value)
		}

	} else {
		sortedLines = listOfStrings
	}

	sort.Strings(sortedLines)

	if parameters.isReverse && parameters.columnNumber == 0 {
		sort.Sort(sort.Reverse(sort.StringSlice(sortedLines)))
	}

	if parameters.columnNumber > 0 {
		keys := make([]string, 0, len(sortedLines))
		m := make(map[string][]string)
		for _, line := range sortedLines {
			currentLine := strings.Split(line, " ")
			key := currentLine[parameters.columnNumber-1]
			keys = append(keys, key)

			if _, ok := m[key]; ok {
				m[key] = append(m[key], line)
			} else {
				m[key] = append(m[key], line)
			}
		}

		keys = removeDuplicateStr(keys)
		sort.Strings(keys)

		for _, key := range keys {
			result = append(result, m[key]...)
		}
		if parameters.isReverse {
			sort.Sort(sort.Reverse(sort.StringSlice(result)))
		}

		return
	}

	if parameters.isNumericSort {
		var keys []float64
		numKeyForString := make(map[float64][]string)

		for _, line := range sortedLines {
			currentLine := strings.Split(line, " ")

			key, err := strconv.ParseFloat(currentLine[0], 64)
			if err != nil {
				return nil, err
			}
			keys = append(keys, key)

			if _, ok := numKeyForString[key]; ok {
				numKeyForString[key] = append(numKeyForString[key], line)
			} else {
				numKeyForString[key] = append(numKeyForString[key], line)
			}
		}

		keys = removeDuplicateInt(keys)

		if parameters.isReverse {
			sort.Sort(sort.Reverse(sort.Float64Slice(keys)))
		} else {
			sort.Sort(sort.Float64Slice(keys))
		}

		for _, key := range keys {
			result = append(result, numKeyForString[key]...)
		}

		return
	}

	result = sortedLines

	return
}

func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	var list []string
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
func removeDuplicateInt(intSlice []float64) []float64 {
	allKeys := make(map[float64]bool)
	var list []float64
	for _, item := range intSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
