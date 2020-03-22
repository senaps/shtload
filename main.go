package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"shtload/utils"
)

func getUrl(url string, data string) ([]byte, string) {
	fullUrl := url + "/"
	resp, err := http.Get(fullUrl)
	if err != nil {
		log.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	return body, resp.Status
}

func benchmarkUrl(url string, method string, data string) {
	for {
		status := ""
		success := 0
		failed := 0
		start := time.Now()
		for {
			if method == utils.GET {
				_, status = getUrl(url, "")
			} else if method == utils.POST {
				_, status = getUrl(url, data)
			}
			finish := time.Now()
			elapsed := finish.Sub(start).Seconds()
			if status == "200 OK" {
				success += 1
			} else {
				failed += 1
			}
			if elapsed >= 1 {
				break
			}
		}
		log.Println("success: ", success, "failed: ", failed)
	}
}

func main() {
	fmt.Println("here we go!...")
	confs := utils.ReadConfig()
	base_url := confs.BaseUrl
	for _, url := range confs.Urls {
		fmt.Println("base_url is: ", base_url)
		fmt.Println("url is: ", url)
		// 	if value == "get" {
		// 		defer benchmarkUrl(value, utils.GET, "")
		// 	}
	}

	// defer benchmarkUrl()
}
