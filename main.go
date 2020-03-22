package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"shtload/utils"
)

var BASE_URL string

func get_url(url string, data string) ([]byte, string) {
	full_url := BASE_URL + url
	resp, err := http.Get(full_url)
	if err != nil {
		log.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	return body, resp.Status
}

func benchmark_url(name string, url string, method string, data string) {
	for {
		status := ""
		success := 0
		failed := 0
		start := time.Now()
		for {
			if method == utils.GET {
				_, status = get_url(url, "")
			} else if method == utils.POST {
				_, status = get_url(url, data)
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
		log.Println("success:", success, "failed:", failed, "url:", name)
	}
}

func main() {
	log.Println("starting the shtload...")
	confs := utils.ReadConfig()
	BASE_URL = confs.Base_Url
	for _, url := range confs.Urls {
		if url.Method == "get" {
			defer benchmark_url(url.Name, url.Route, utils.GET, "")
		}
	}

	// defer benchmark_url()
}
