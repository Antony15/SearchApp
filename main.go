package main

import (
	"fmt"
	"os"
	"strings"
	colour "github.com/fatih/color"
)

func main() {
	args 		:= strings.Join(os.Args[1:], " ")
	searching 	:= strings.Replace(args, ",", " & ", -1)
	searches 	:= strings.Split(args, ",")
	fmt.Printf("Searching for: %s\n", searching)
	fmt.Println("")
	for _, s := range searches {
		result := getSearchs(s)
		colour.Cyan(`Heading:	%s`, result.Heading)
		colour.Blue(`AbstractURL:	%s`, result.AbstractURL)
		colour.Green(`AbstractText:	%s`, result.AbstractText)
		fmt.Println("")
	}

}
