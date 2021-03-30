package main

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

func main() {
	chanelSampleFunc()
}

func NonBufferSample() {
	size := 2
	ch := make(chan string, size)
	send(ch, "one")
	send(ch, "two")
	go send(ch, "three")
	go send(ch, "four")

	fmt.Println("All data sent to the channel")

	for i := 0; i < 4; i++ {
		fmt.Println(<-ch)
	}

	fmt.Println("Done!")
}

func send(ch chan string, message string) {
	ch <- message
}

func chanelSampleFunc() {
	proxyUrl, _ := url.Parse("http://localhost:8888")

	http.DefaultTransport = &http.Transport{
		Proxy: http.ProxyURL(proxyUrl),
	}

	start := time.Now()

	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com/",
		"https://api.somewhereintheinternet.com/",
		"https://graph.microsoft.com",
	}
	ch := make(chan string, 10)
	for _, api := range apis {
		go checkAPIChannel(api, ch)
	}

	for i := 0; i < len(apis); i++ {
		fmt.Print(<-ch)
	}

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}

func checkAPINonChannel(api string) {
	_, err := http.Get(api)

	if err != nil {
		fmt.Printf("ERROR: %s is down!\n", api)
		return
	}

	fmt.Printf("SUCCESS: %s is up and running!\n", api)
}

func checkAPIChannel(api string, ch chan string) {
	_, err := http.Get(api)

	if err != nil {
		ch <- fmt.Sprintf("ERROR: %s is down!\n", api)
		return
	}

	ch <- fmt.Sprintf("SUCCESS: %s is up and running!\n", api)
}
