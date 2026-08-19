package main

import (
	"log"
	"net/http"
)

func main() {
	testGetURL()
}

func testGetURL() {
	dst := getRequestDestination()
	for _, url := range dst {
		status := getReqStatus(url)
		println(url, "\n\t", status, "\n")
	}
}

func getRequestDestination() []string {
	r := []string{}
	rU := getURL()
	rP := getPATH()

	for _, u := range rU {
		for _, p := range rP {
			r = append(r, "https://"+u+"/"+p)
		}
	}

	return r
}

func getURL() []string {
	return []string{"google.com", "ya.ru", "bing.com"}
}

func getPATH() []string {
	return []string{"index.html", "robot.txt", "favicon.ico"}
}

func getReqStatus(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	return resp.Status
}
