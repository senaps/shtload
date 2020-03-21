package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"shtload/utils"
)

func get_url(url string, data string) ([]byte, string) {
	full_url := url + "/"
	resp, err := http.Get(full_url)
	if err != nil {
		log.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	return body, resp.Status
}

func benchmark_url(url string, method string, data string) {
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
		log.Println("success: ", success, "failed: ", failed)
	}
}

func main() {
	fmt.Println("here we go!...")
	confs := utils.ReadConfig()
	for key, value := range confs {
		if value == "get" {
			defer benchmark_url(value, utils.GET, "")
		}
	}

	// defer benchmark_url()
}
