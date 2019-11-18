package main

import (
	"encoding/json"
	"os"
	"time"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/karlseguin/ccache"
	colour "github.com/fatih/color"
)

var cache 		= ccache.New(ccache.Configure().MaxSize(5))

const (
	apiURL       = "https://api.duckduckgo.com/"
)

type Search struct {
	Heading			string			`json:"Heading"`
	AbstractURL     string         	`json:"AbstractURL"`
	AbstractText    string      	`json:"AbstractText"`
}

func getSearchs(name string) Search {	
	var search Search
	item := cache.Get(name)
	log.Println(item)
	if item == nil {		
		resp, err := http.Get(apiURL + "?q="+name+"&format=json")
		if err != nil {
			log.Fatalf("Error retrieving data: %s\n", err)
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("Error reading data: %s\n", err)
		}
		cache.Set(name, body, time.Minute * 10)
		json.Unmarshal(body, &search)

		if (Search{}) == search {

			colour.Red("Sorry the Search doesn't exist")
			os.Exit(1)
			return search
		}
	}
	return search
}
