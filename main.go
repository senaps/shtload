package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var BASE_URL = os.Getenv("BASE_URL")

func get_url(url string, data string) ([]byte, string) {
	full_url := BASE_URL + url + "/"
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
			if method == "get" {
				_, status = get_url(url, "")
			} else if method == "post" {
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
	defer benchmark_url("ping", "get", "")
	// defer benchmark_url()
}
