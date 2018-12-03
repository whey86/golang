package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func importTextFile(textFile string) string {
	b, err := ioutil.ReadFile(textFile)
	if err != nil {
		fmt.Print(err)
	}
	return string(b)
}

func buildCommandList (string input) command[] {
	var commands[]command;
	var scanner Scanner= bufio.NewScanner(strings.NewReader(input))
	re := regexp.MustCompile("(\\d+)")
	for scanner.Scan() {
		var line string = scanner.Text();
		var match[]string = re.FindAllString(line, -1);
		commands := append(commands, commad{start[0]:match[1]});
	}
	return commands;
}
type command struct {
	 start [2]int;
	 end [2]int;
	 sign string;
}

func main() {
	fmt.Println("Advent of code day 1 part 1")
	var sum int = 0
	var frq []int
	input := importTextFile("first.txt")

	re := regexp.MustCompile("[0-9]+")

	scanner := bufio.NewScanner(strings.NewReader(input))

	for scanner.Scan() {
		var myString = scanner.Text()
		//fmt.Println(myString);
		frq = append(frq, sum)
		if strings.Contains(myString, "+") {

			testmatch := re.FindStringSubmatch(myString)
			extractInt, err := strconv.Atoi(testmatch[0])
			if err != nil {
				// handle error
				fmt.Println(err)
			}
			sum += extractInt

		} else {
			testmatch := re.FindStringSubmatch(myString)
			extractInt, err := strconv.Atoi(testmatch[0])
			if err != nil {
				// handle error
				fmt.Println(err)
			}
			sum -= extractInt

		}

	}
	fmt.Println("Frequecy is: ", sum)
}
