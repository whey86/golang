package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"strings"
)

func importTextFile(textFile string) string {
	b, err := ioutil.ReadFile(textFile)
	if err != nil {
		fmt.Print(err)
	}
	return string(b)
}

func getStringArray() []string {
	input := importTextFile("first.txt")
	var arrayOflines []string
	scanner := bufio.NewScanner(strings.NewReader(input))

	for scanner.Scan() {
		var line = scanner.Text()
		arrayOflines = append(arrayOflines, line)
	}
	return arrayOflines

}
func main() {
	fmt.Println("Advent of code day 1 part 1")

	var lines []string = getStringArray()
	var counter int = 0
	for i := 0; i < len(lines); i++ {

		for j := 0; j < len(lines); j++ {
			if i == j {
				continue
			}
			var mainLine []string = strings.Split(lines[i], "")
			var compareLine []string = strings.Split(lines[j], "")
			for k := 0; k < len(mainLine); k++ {
				if mainLine[k] != compareLine[k] {
					counter++
				}
			}
			if counter == 1 {
				fmt.Println(lines[i])
				fmt.Println(lines[j])
			}
			counter = 0
		}
	}

}
