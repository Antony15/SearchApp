 
package main
 
import (
"fmt"
"time"
"net/http"
"io/ioutil"
"encoding/json"
"github.com/karlseguin/ccache"
)

var cache = ccache.New(ccache.Configure().MaxSize(5))

func main() {

	var data map[string]interface{}
	item := cache.Get("tata motors")
	fmt.Println(item)
	if item == nil {
		url := "https://api.duckduckgo.com/?q=tata motors&format=json"
		 
		req, _ := http.NewRequest("GET", url, nil)
		 
		res, _ := http.DefaultClient.Do(req)
		 
		defer res.Body.Close()

		body, _ := ioutil.ReadAll(res.Body)

		if err := json.Unmarshal(body, &data); err != nil {
			panic(err)
		}
		cache.Set("tata motors", body, time.Minute * 10)
		
		fmt.Println(data["Heading"])
		fmt.Println(data["AbstractURL"])
		fmt.Println(data["AbstractText"])
	} else {
	  //~ data := item.Value().(*User)
	  fmt.Println("no")
	}	 
	

}
