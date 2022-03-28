package main

import "time"

type TickersStruct struct {
	Results []struct {
		Ticker          string    `json:"ticker"`
		Name            string    `json:"name"`
		Market          string    `json:"market"`
		Locale          string    `json:"locale"`
		PrimaryExchange string    `json:"primary_exchange"`
		Type            string    `json:"type,omitempty"`
		Active          bool      `json:"active"`
		CurrencyName    string    `json:"currency_name"`
		Cik             string    `json:"cik,omitempty"`
		CompositeFigi   string    `json:"composite_figi,omitempty"`
		ShareClassFigi  string    `json:"share_class_figi,omitempty"`
		LastUpdatedUtc  time.Time `json:"last_updated_utc"`
	} `json:"results"`
	Status    string `json:"status"`
	RequestID string `json:"request_id"`
	Count     int    `json:"count"`
	NextURL   string `json:"next_url"`
}
