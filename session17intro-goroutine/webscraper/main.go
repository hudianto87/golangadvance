package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type SharedData struct {
	mu   sync.Mutex
	data map[string]float64
}

func scrapeWebSite(url string, shared *SharedData, wg *sync.WaitGroup, r *rand.Rand) {
	defer wg.Done()

	time.Sleep(time.Duration(r.Intn(1000)) * time.Millisecond)

	scrapedData := r.Float64() * 1000

	shared.mu.Lock()
	shared.data[url] = scrapedData
	shared.mu.Unlock()

	fmt.Println("scrape data from %s: %f\n", url, scrapedData)
}

func main() {
	source := rand.NewSource(time.Now().UnixNano())

	r := rand.New(source)

	SharedData := &SharedData{
		data: make(map[string]float64),
	}

	website := []string{
		"https://finance.yahoo.com",
		"https://www.investing.com",
		"https://www.alphavantage.com",
		"https://www.google.com/finance",
		"https://www.nasdaq.com",
		"https://www.bloomberg.com",
		"https://www.morningstart.com",
		"https://coinmarketcap.com",
		"https://data.worldbank.org",
		"https://www.quandl.com",
	}

	var wg sync.WaitGroup

	for _, url := range website {
		wg.Add(1)
		go scrapeWebSite(url, SharedData, &wg, r)
	}

	wg.Wait()

	fmt.Println("collect financial data ")
	//SharedData.mu.Lock()
	for site, value := range SharedData.data {
		fmt.Printf("%s: %f\n", site, value)
	}

	//SharedData.mu.Unlock()
}
