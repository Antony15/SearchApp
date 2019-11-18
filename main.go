package main

import (
	"os"
	"github.com/spf13/cobra"
	"strings"
	colour "github.com/fatih/color"
)

func main() {
	cmd := &cobra.Command{
		Use:          "SearchApp",
		Short:        "CLI based search application",
		RunE: func(cmd *cobra.Command, args []string) error {
			argsearch 	:= strings.Join(args, " ")
			searching 	:= strings.Replace(argsearch, ",", " & ", -1)
			searches 	:= strings.Split(argsearch, ",")
			cmd.Println(argsearch)
			cmd.Printf("Searching for: %s\n", searching)
			cmd.Println("")
			for _, s := range searches {
				result := getSearchs(s)
				colour.Cyan(`Heading:	%s`, result.Heading)
				colour.Blue(`AbstractURL:	%s`, result.AbstractURL)
				colour.Green(`AbstractText:	%s`, result.AbstractText)
				cmd.Println("")
			}
			return nil
		},
	}


	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
