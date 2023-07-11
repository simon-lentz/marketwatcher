package main

import (
	"testing"
)

func TestTickers_GetPrices(t *testing.T) {
	s := Stock{
		Client: client,
	}

	p, err := s.GetPrices()
	if err != nil {
		t.Error(err)
	}

	expected := 0.00
	if p.Current != expected {
		t.Error("Ticker price did not match expected value:", p.Current)
	}
}
