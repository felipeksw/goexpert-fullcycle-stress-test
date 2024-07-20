package main

import (
	"flag"
	"fmt"
	"log/slog"

	"github.com/felipeksw/goexpert-fullcycle-stress-test/internal/usecase"
)

func main() {

	var (
		url         string
		requests    int
		concurrency int
	)

	flag.StringVar(&url, "url", "", "URL of the web service to be tested")
	flag.IntVar(&requests, "requests", 0, "Total number of requests to be made")
	flag.IntVar(&concurrency, "concurrency", 1, "Number of simultaneous calls")

	flag.Parse()

	st := usecase.NewStressTest(url, requests, concurrency)

	ret, err := st.Execute()
	if err != nil {
		slog.Error("[Execute]", "error", err.Error())
	}

	fmt.Println("Total of execution time:", ret.TotalTime)
	fmt.Println("Total requests:", ret.TotalRequests)
	fmt.Println("Total HTTP status 200:", ret.TotalSuccessRequests)
	for k, v := range ret.ErrorRequests {
		fmt.Printf("Total HTTP status %d: %d\n", k, v)
	}
}
