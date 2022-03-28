package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const basePath = "https://api.polygon.io"

//GetFromUrlAndClean is a function that gets data from a url, cleans it and returns it as a TickersStruct
func GetFromUrlAndClean(url string, apiKey string) TickersStruct {
	tickersStruct := TickersStruct{}
	url = fmt.Sprintf("%s&apiKey=%s", url, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, &tickersStruct)
	if err != nil {
		panic(err)
	}
	return tickersStruct
}

//GetAllTickerNames returns all the ticker names
func GetAllTickerNames(apiKey string, results *[]TickersStruct, m model) {
	startUrl := fmt.Sprintf("%s/v3/reference/tickers?active=true&sort=ticker&order=asc&limit=1000", basePath)
	tickers := GetFromUrlAndClean(startUrl, apiKey)
	for tickers.NextURL != "" {
		tickers := GetFromUrlAndClean(tickers.NextURL, apiKey)
		*results = append(*results, tickers)
	}
}

func GetAllTickerDetails(apiKey string, results *[]TickersStruct) {
	startUrl := fmt.Sprintf("%s/v3/reference/tickers?active=true&sort=ticker&order=asc&limit=1000", basePath)
	tickers := GetFromUrlAndClean(startUrl, apiKey)
	for tickers.NextURL != "" {
		tickers := GetFromUrlAndClean(tickers.NextURL, apiKey)
		*results = append(*results, tickers)
	}
}
