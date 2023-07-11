package main

import (
	"log"
	"testing"
)

func TestApp_getTickerText(t *testing.T) {

	currentPrice, _ := testApp.getTickerText()
	log.Println(currentPrice.Text)
	if currentPrice.Text == "Current Price: $0.0000" { // fix tests with real values when the API issues are figured out
		t.Error("current price does not match expected value")
	}

}
