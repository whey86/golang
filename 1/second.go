package main

import (
	"fmt"
	"io/ioutil"
	"bufio"
	"strings"
	"regexp"
	"strconv"
)

func importTextFile(textFile string) string{
	b, err := ioutil.ReadFile(textFile);
	if err != nil {
        fmt.Print(err)
    }
	return string(b);
}



func main() {
	fmt.Println("Advent of code day 1 part 1");
	var sum int = 0;
	var frq [] int;
	for true {
	input := importTextFile("first.txt");


	re := regexp.MustCompile("[0-9]+");

	scanner := bufio.NewScanner(strings.NewReader(input))


	for scanner.Scan() {
		var myString = scanner.Text();
		//fmt.Println(myString);
		frq = append(frq,sum);
		if strings.Contains(myString,"+"){
		
			testmatch := re.FindStringSubmatch(myString);
			 extractInt,err := strconv.Atoi(testmatch[0]);
			    if err != nil {
        // handle error
        fmt.Println(err)
				}
		sum += extractInt
      
    }else{
					testmatch := re.FindStringSubmatch(myString);
			 extractInt,err := strconv.Atoi(testmatch[0]);
			    if err != nil {
        // handle error
        fmt.Println(err)
				}
		sum -= extractInt

	}
			for i := 0; i < len(frq); i++ {
				
 //  fmt.Println(frq[i], " ---- ",sum );

					if frq[i] == sum {
						fmt.Println("first double frq is ", sum);
						return;
					}
	}
	}	
		
}
	 fmt.Println("iteration");
}


