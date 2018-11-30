package main

import (
	"fmt"
	"io/ioutil"
)


func importTextFile(textFile string) string{
	b, err := ioutil.ReadFile(textFile);
	if err != nil {
        fmt.Print(err)
    }
	return string(b);
}

func main() {
	fmt.Println("Dzien dobre kochanie");

}
