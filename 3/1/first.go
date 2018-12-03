package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var pattern [][]string
var TAKEN string = "@"
var OVERRITTEN = "X"
var EMPTY string = "."
var takenCounter int = 0

func importTextFile(textFile string) string {
	b, err := ioutil.ReadFile(textFile)
	if err != nil {
		fmt.Print(err)
	}
	return string(b)
}

func getStringArray() []string {
	input := importTextFile("test.txt")
	var arrayOflines []string
	scanner := bufio.NewScanner(strings.NewReader(input))

	for scanner.Scan() {
		var line = scanner.Text()
		arrayOflines = append(arrayOflines, line)
	}
	return arrayOflines

}

func buildCommandList(input []string) []command {
	var commands []command

	for i := 0; i < len(input); i++ {
		re := regexp.MustCompile("(\\d+)")
		var line string = input[i]
		var match []string = re.FindAllString(line, -1)
		ex, _ := strconv.Atoi(match[3])
		sx, _ := strconv.Atoi(match[1])
		ey, _ := strconv.Atoi(match[4])
		sy, _ := strconv.Atoi(match[2])
		commands = append(commands, command{
			startX: sx,
			startY: sy,
			endX:   ex + sx,
			endY:   ey + sy})
	}

	return commands
}

func addToPattern(x int, y int) {
	horizontalSize := len(pattern)

	//if array does not exist
	if horizontalSize < x {
		for i := horizontalSize; i <= x; i++ {
			pattern = append(pattern, []string{})
			for j := 0; j <= y; j++ {
				pattern[i] = append(pattern[i], ".")
			}
		}
	}
	currentPattern := pattern[x][y]
	switch p := currentPattern; p {
	case EMPTY:
		pattern[x][y] = TAKEN
	case TAKEN:
		pattern[x][y] = OVERRITTEN
		takenCounter++
	}

}

type command struct {
	startX int
	endX   int
	startY int
	endY   int
	sign   string
}

func main() {
	fmt.Println("Advent of code day 3 part 1")
	commands := buildCommandList(getStringArray())

	for _, element := range commands {

		for i := element.endX; i >= element.startX; i-- {
			for j := element.endY; j >= element.startY; j-- {
				addToPattern(i, j)
			}
		}
	}

	fmt.Println("Number of overritten ", takenCounter)
}
