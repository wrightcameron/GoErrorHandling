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
	//First Scanner breaks text into lines
	lineScanner := bufio.NewScanner(strings.NewReader(text))
	lineScanner.Split(bufio.ScanLines)
	for lineScanner.Scan() {
		wordScanner := bufio.NewScanner(strings.NewReader(lineScanner.Text()))
		wordScanner.Split(bufio.ScanWords)
		for wordScanner.Scan() {
			word := wordScanner.Text()
			if isUpperCase(word) {
				color.Red(word)
			} else {
				color.Blue(word)
			}
		}
		//fmt.Println()
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
