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

// Initialize ccache with maximum size of 5
var cache 		= ccache.New(ccache.Configure().MaxSize(5))

// Constant Api URL
const (
	apiURL       = "https://api.duckduckgo.com/"
)

// Search struct
type Search struct {
	Heading			string			`json:"Heading"`
	AbstractURL     string         	`json:"AbstractURL"`
	AbstractText    string      	`json:"AbstractText"`
}

// Function for getting searches
func getSearchs(name string) Search {	
	var search Search
	
	// Get search value from cache
	item := cache.Get(name)	
	
	// if cache nil
	if item == nil {
		// Make a cURL get request 
		resp, err := http.Get(apiURL + "?q="+name+"&format=json")
		
		// Log error when request not successful 
		if err != nil {
			log.Fatalf("Error retrieving data: %s\n", err)
		}
		
		// Close response body
		defer resp.Body.Close()	
		
		// Read response body		
		body, err := ioutil.ReadAll(resp.Body)
		
		// Log error when response body could not be read
		if err != nil {
			log.Fatalf("Error reading data: %s\n", err)
		}
		
		// Set response in cache with search name as key
		cache.Set(name, body, time.Minute * 10)
		
		// Unmarshall json response body to search struct
		json.Unmarshal(body, &search)

	}else{
		// Unmarshal json in cache to search struct
		if err := json.Unmarshal(item.Value().([]byte), &search); err != nil {
			panic(err)
		}
	}
	// If search struct is nil
	if (Search{}) == search {
		colour.Red("Could not find any result")
		os.Exit(1)
		return search
	}
	
	// Return search struct	
	return search
}
