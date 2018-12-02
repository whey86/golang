package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func importTextFile(textFile string) string {
	b, err := ioutil.ReadFile(textFile)
	if err != nil {
		fmt.Print(err)
	}
	return string(b)
}

func SortString(w string) []string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return s
}

func main() {
	var testString string = "bcaa"
	fmt.Println(SortString(testString))
	fmt.Println("Advent of code day 1 part 1")
	var countTwice int = 0
	var countTrice int = 0

	input := importTextFile("first.txt")

	scanner := bufio.NewScanner(strings.NewReader(input))
	//For each line
	for scanner.Scan() {
		var line = scanner.Text()
		var sortedLine []string = SortString(line)
		var lastCharater string = ""
		var counter int = 1

		var nTwo = false
		var nThree = false
		// For each character
		fmt.Println(sortedLine)
		for i := 0; i < len(sortedLine); i++ {
			if lastCharater == sortedLine[i] {
				counter++
			}
			if lastCharater != sortedLine[i] || i == len(sortedLine)-1 {
				if counter == 2 {
					fmt.Println("Added twice", counter)
					counter = 1
					nTwo = true
				}
				if counter == 3 {
					fmt.Println("Added trice", counter)
					counter = 1
					nThree = true
				}
			}
			lastCharater = sortedLine[i]
		}
		if nTwo {
			countTwice++
		}
		if nThree {
			countTrice++
		}

	}
	fmt.Println("Checksum is: ", countTwice*countTrice)
}
