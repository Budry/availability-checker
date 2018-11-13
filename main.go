package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"

	"bitbucket.org/Budry/availability-checker/src/email"
	"bitbucket.org/Budry/availability-checker/src/options"
	"bitbucket.org/Budry/availability-checker/src/sites"
)

func main() {
	file, err := os.Open("/var/lib/availability-checker/test.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	configOptions := &options.Options{}
	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	json.Unmarshal([]byte(byteValue), &configOptions)

	var wg sync.WaitGroup
	var results []*sites.Result
	wg.Add(len(configOptions.Sites))
	for _, site := range configOptions.Sites {
		go func(site sites.Site) {
			defer wg.Done()
			result := site.Process()
			results = append(results, site.Process())
			if result.HasError() {
				email.SendFailNotificationMessage(result)
			}
		}(site)
	}
	wg.Wait()
	for _, result := range results {
		if result.HasError() {
			fmt.Print("[FAILED] ")
		} else {
			fmt.Print("[SUCCESS] ")
		}
		fmt.Print(result.Site.Url)
		if result.HasError() {
			fmt.Print(":\n")
			for _, errorMessage := range result.Errors {
				fmt.Println("\tError: " + errorMessage)
			}
		} else {
			fmt.Println()
		}
	}
}
