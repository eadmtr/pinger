package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	c "pinger/client"
)

type URLData struct {
	ursStr string
	status string
}

type URLDataResults struct {
	sync.Mutex
	results []URLData
}

func main() {
	r := getURLWg()

	for _, ud := range r {
		fmt.Println(ud.ursStr, ud.status)
	}
}

func getURLWg() []URLData {
	r := URLDataResults{results: make([]URLData, 0)}
	dst := c.GetRequestDestination()

	wg := &sync.WaitGroup{}
	wg.Add(len(dst))

	for _, url := range dst {
		ud := URLData{url, ""}

		go func() {
			defer wg.Done()
			ud, err := getURLStatus(ud)

			if err == nil {
				r.Lock()
				r.results = append(r.results, ud)
				r.Unlock()
			}
		}()
	}
	wg.Wait()

	return r.results
}

func getURLStatus(urlData URLData) (URLData, error) {
	resp, err := http.Get(urlData.ursStr)
	if err != nil {
		log.Fatalln(err)
		return urlData, err
	}

	urlData.status = resp.Status
	return urlData, nil
}
