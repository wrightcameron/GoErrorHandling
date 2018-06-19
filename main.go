package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/fatih/color"
)

func main() {
	data := readTxtFile("Gettysburg-Address.txt")
	findUpperCaseWords(data)
}

func findUpperCaseWords(text string) {
	scanner := bufio.NewScanner(strings.NewReader(text))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		if isUpperCase(word) {
			color.Red(word)
		} else {
			color.Blue(word)
		}
	}
}

func isUpperCase(word string) bool {
	if word[0] >= 65 && word[0] <= 90 {
		return true
	}
	return false
}

func readTxtFile(fileName string) string {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Can't read file, or does not exist", fileName)
		panic(err)
	}
	return string(data)
}
