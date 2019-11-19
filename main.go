package main

import (
	"os"
	"fmt"
	"strings"
	colour "github.com/fatih/color"
)

func main() {
	// Get search request arguments
	argsearch 	:= strings.Join(os.Args[1:], " ")
	searching 	:= strings.Replace(argsearch, ",", " & ", -1)
	searches 	:= strings.Split(argsearch, ",")
	// If search not null
	if len(os.Args)>1{
		fmt.Printf("Searching for: %s\n", searching)
		fmt.Println("")	
		
		// Loop for multiple searches			
		for _, s := range searches {
			
			// Get search results
			result := getSearchs(s)
			
			// Use colour package to highlight
			colour.Cyan(`Heading:	%s`, result.Heading)
			colour.Blue(`AbstractURL:	%s`, result.AbstractURL)
			colour.Green(`AbstractText:	%s`, result.AbstractText)
			fmt.Println("")
		}
	}else{
		colour.Red("Search cannot be null")
		os.Exit(1)				
	}
}
