package usecase

import (
	"errors"
	"net/http"
	"sync"
	"time"
)

type StressTestResultDto struct {
	TotalTime            time.Duration
	TotalRequests        int
	TotalSuccessRequests int
	ErrorRequests        map[int]int
}

type StressTest struct {
	Url         string
	Requests    int
	Concurrency int
}

func NewStressTest(url string, requests int, concurrency int) *StressTest {
	return &StressTest{
		Url:         url,
		Requests:    requests,
		Concurrency: concurrency,
	}
}

func (s *StressTest) IsValid() error {
	msg := ""
	if len(s.Url) < 3 {
		msg += "invalid url; "
	}

	if s.Requests < s.Concurrency {
		msg += "invalid request or concurrency; "
	}

	if s.Requests == 0 {
		msg += "invalid requests; "
	}

	if s.Concurrency == 0 {
		msg += "invalid concurrency; "
	}

	if msg != "" {
		return errors.New(msg)
	}

	return nil
}

func (s *StressTest) Execute() (*StressTestResultDto, error) {

	err := s.IsValid()
	if err != nil {
		return nil, err
	}

	start := time.Now()
	var statusErrors map[int]int = map[int]int{}
	totalSuccessRequests := 0
	totalErrorsRequests := 0
	requests := s.Requests
	concurrency := s.Concurrency
	waitGroup := sync.WaitGroup{}

	chSt := make(chan int, concurrency)
	go func() {
		for x := range chSt {
			if x == 200 {
				totalSuccessRequests++
				waitGroup.Done()
				continue
			}
			totalErrorsRequests++
			statusError, ok := statusErrors[x]
			if !ok {
				statusErrors[x] = 1
				waitGroup.Done()
				continue
			}
			statusErrors[x] = statusError + 1
			waitGroup.Done()
		}
		close(chSt)
	}()

	executed := 0
	for {
		concurrency := concurrency
		if (requests - executed) < concurrency {
			concurrency = requests - executed
		}
		if executed >= requests {
			break
		}
		waitGroup.Add(concurrency)
		for i := 0; i < concurrency; i++ {
			go func() {
				resp, _ := http.DefaultClient.Get(s.Url)
				chSt <- resp.StatusCode
			}()
		}
		waitGroup.Wait()
		executed += concurrency

	}

	return &StressTestResultDto{
		TotalTime:            time.Now().Sub(start),
		TotalRequests:        totalSuccessRequests + totalErrorsRequests,
		TotalSuccessRequests: totalSuccessRequests,
		ErrorRequests:        statusErrors,
	}, nil
}
