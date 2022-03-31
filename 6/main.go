package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
)

func delimBytes(lines [][]byte, delimByte []byte, fields *int) {
	res := make([][][]byte, 0)
	for i := range lines {
		temp := bytes.Split(lines[i], delimByte)
		if field := *fields; field > -1 {
			if len(temp) > field {
				temp = [][]byte{temp[field]}
			} else {
				temp = [][]byte{}
			}
		}
		res = append(res, temp)
	}

	fmt.Printf("%q", res)
}

func separatedBytes(lines [][]byte, delimByte []byte) [][]byte {
	temp := make([][]byte, 0)
	for i := range lines {
		if bytes.Contains(lines[i], delimByte) {
			temp = append(temp, lines[i])
		}
	}
	lines = temp
	return lines
}

func main() {

	var fields = flag.Int("f", -1, "выбрать колонки")
	var delim = flag.String("d", " ", "использовать другой разделитель")
	var separated = flag.Bool("s", false, "обрабатывать только строки с разделителем")

	flag.Parse()

	nameFile := flag.Arg(0)

	content, err := ioutil.ReadFile(nameFile)
	if err != nil {
		panic(err)
	}

	delimByte := []byte(*delim)
	lines := bytes.Split(content, []byte("\n"))

	if *separated {
		separatedBytes(lines, delimByte)
	}

	delimBytes(lines, delimByte, fields)

}
