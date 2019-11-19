package main

import (
	"os"
	"github.com/spf13/cobra"
	"strings"
	colour "github.com/fatih/color"
)

func main() {
	// Initialize Cobra CLI
	cmd := &cobra.Command{
		Use:          "SearchApp",
		Short:        "CLI based search application",
		RunE: func(cmd *cobra.Command, args []string) error {
			argsearch 	:= strings.Join(args, " ")
			searching 	:= strings.Replace(argsearch, ",", " & ", -1)
			searches 	:= strings.Split(argsearch, ",")
			// If search not null
			if len(args)>0{
				cmd.Printf("Searching for: %s\n", searching)
				cmd.Println("")	
				
				// Loop for multiple searches			
				for _, s := range searches {
					
					// Get search results
					result := getSearchs(s)
					
					// Use colour package to highlight
					colour.Cyan(`Heading:	%s`, result.Heading)
					colour.Blue(`AbstractURL:	%s`, result.AbstractURL)
					colour.Green(`AbstractText:	%s`, result.AbstractText)
					cmd.Println("")
				}
			}else{
				colour.Red("Search cannot be null")
				os.Exit(1)				
			}
			return nil
		},
	}
	
	// Execute Cobra CLI
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
