package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var pattern [][]tile
var TAKEN string = "@"
var OVERRITTEN = "X"
var EMPTY string = "."
var takenCounter int = 0
var isChanged bool = false
var clean []int

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

func addToColumn(x int, y int, index int) {
	//add empty to row
	if len(pattern[x]) <= y {
		for i := len(pattern[x]); i <= y; i++ {
			pattern[x] = append(pattern[x], tile{sign: EMPTY})
		}
	}
	addPattern(x, y, index)
}

func addToRow(x int, y int, index int) {
	horizontalSize := len(pattern)

	//if array does not exist
	if horizontalSize <= x {
		for i := horizontalSize; i <= x; i++ {
			pattern = append(pattern, []tile{})
		}
	}
	addToColumn(x, y, index)

}
func addPattern(x int, y int, index int) {

	pattern[x][y].order = append(pattern[x][y].order, index)
	switch p := pattern[x][y].sign; p {
	case EMPTY:
		pattern[x][y].sign = TAKEN
	case TAKEN:
		pattern[x][y].sign = OVERRITTEN
		takenCounter++
		removeAll(pattern[x][y].order)
	case OVERRITTEN:
		removeAll(pattern[x][y].order)
	}
}
func removeAll(orders []int) {
	for _, element := range orders {
		clean[element-1] = 0
	}

}

func stringInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func remove(s []int, i int) {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	clean = s[:len(s)-1]
}
func searchUntouched() {
	var first bool = true
	var row int
	var col int
	for i := 0; i < len(pattern); i++ {
		for j := 0; j < len(pattern[i]); j++ {
			if pattern[i][j].sign == TAKEN {
				if first {
					first = false
					fmt.Println("NON OVERLAPPING BOX IS", i, "+", j)
				}
				row = i
				col = j
			}
		}

	}
	fmt.Println("NON OVERLAPPING BOX IS", row, "+", col)
}

type command struct {
	startX int
	endX   int
	startY int
	endY   int
	sign   string
}

type tile struct {
	sign  string
	order []int
}

func main() {
	fmt.Println("Advent of code day 3 part 1")
	commands := buildCommandList(getStringArray())

	for index, element := range commands {
		clean = append(clean, index+1)
		for i := element.endX; i > element.startX; i-- {
			for j := element.endY; j > element.startY; j-- {
				addToRow(i, j, index+1)
			}
		}
		if !isChanged {
			//	fmt.Println("NON OVERLAPPING BOX IS", index+1)

		}
	}

	fmt.Println("Number of overritten ", clean)
	//	searchUntouched()
}
