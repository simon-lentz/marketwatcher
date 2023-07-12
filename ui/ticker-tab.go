// https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=IBM&apikey=292F1EEKKBF9AVOE
package ui

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

var symbol = "IBM" // test value

type Stock struct {
	Symbol string
	Price  Price `json:"Global Test"`
	Client *http.Client
}

type Price struct {
	Current float64   `json:"05. price"`
	Time    time.Time `json:"-"`
}

func (s *Stock) GetPrices() (*Price, error) {
	if s.Client == nil {
		s.Client = &http.Client{}
	}

	client := s.Client

	url := fmt.Sprintf("https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=%s&apikey=292F1EEKKBF9AVOE", symbol)

	req, _ := http.NewRequest("GET", url, nil)

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("error contacting alphavantage: %v", err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("error parsing response body: %v", err)
		return nil, err
	}

	stock := Stock{
		Symbol: symbol,
	}

	if err = json.Unmarshal(body, &stock); err != nil {
		log.Printf("error parsing json: %v", err)
		return nil, err
	}

	current := stock.Price.Current

	var currentInfo = Price{
		Current: current,
		Time:    time.Now(),
	}

	return &currentInfo, nil
}
