// Package client
package client

import (
	"log"
	"net/http"
)

type URLData struct {
	ursStr string
	status string
}

func getDomain() []string {
	return []string{"google.com", "ya.ru", "mail.ru"}
}

func getPATH() []string {
	return []string{"/", "/index.html", "/robot.txt", "/favicon.ico"}
}

func GetRequestDestination() []string {
	r := []string{}
	rU := getDomain()
	rP := getPATH()

	for _, u := range rU {
		for _, p := range rP {
			r = append(r, "https://"+u+p)
		}
	}

	return r
}

func URLStatus(urlData URLData) (URLData, error) {
	resp, err := http.Get(urlData.ursStr)
	if err != nil {
		log.Fatalln(err)
		return urlData, err
	}

	urlData.status = resp.Status
	return urlData, nil
}
