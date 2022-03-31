package main

import (
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

func main() {
	after := flag.Int("A", 0, "after key")
	before := flag.Int("B", 0, "before key")
	context := flag.Int("C", 0, "context key")

	count := flag.Int("c", 0, "count key")

	ignoreCase := flag.Bool("i", false, "ignore case key")
	invert := flag.Bool("v", false, "invert key")
	fixed := flag.Bool("F", false, "fixed key")
	lineNum := flag.Bool("n", false, "live num key")
	flag.Parse()

	lines := struct {
		left  int
		right int
		count int
	}{left: 0, right: 0, count: 0}

	find := flag.Arg(0)
	fileName := flag.Arg(1)

	sliceFile := handleFile(fileName)

	if *count != 0 {
		lines.count = *count
	} else {
		lines.count = len(sliceFile)
	}

	switch {
	case *context != 0:
		lines.left, lines.right = *context, *context
	case *after != 0:
		lines.right = *after
	case *before != 0:
		lines.left = *before

	default:
		lines.right = len(sliceFile)
	}

	for i, str := range sliceFile {

		if handleStr(str, find, *ignoreCase, *invert, *fixed) {

			numFirst := math.Max(0, float64(i-lines.left))
			numLast := math.Min(float64(len(sliceFile)-1), float64(i+lines.right))

			for j := numFirst; j <= numLast; j++ {

				if lines.count > 0 {
					if *lineNum {
						fmt.Println("numLine: ", j, "\n", sliceFile[int(j)])
						lines.count--
					} else {
						fmt.Println(sliceFile[int(j)])
						lines.count--
					}
				}

			}
		}

	}

}

func handleFile(nameFile string) []string {
	file, _ := os.Open(nameFile)
	defer file.Close()

	byteFile, _ := io.ReadAll(file)
	stringFile := string(byteFile)
	return strings.Split(stringFile, "\n")
}

func handleStr(str string, findWord string, ignoreCase bool, invert bool, fixed bool) bool {

	if ignoreCase {
		str = strings.ToLower(str)
		findWord = strings.ToLower(findWord)
	}

	if fixed {
		if findWord == strings.Replace(str, "\r", "", 1) {
			return !invert
		}
		return invert

	}

	if strings.Contains(str, findWord) {
		return !invert
	}
	return invert

}
