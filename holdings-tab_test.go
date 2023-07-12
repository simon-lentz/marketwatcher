package main

import "testing"

func TestConfig_getHoldings(t *testing.T) {
	holdings, err := testApp.currentHoldings()
	if err != nil {
		t.Error("failed to retrieve current holding from db:", err)
	}
	if len(holdings) != 2 {
		t.Error("wrong number of holdings")
	}
}

func TestConfig_getHoldingSlice(t *testing.T) {
	slice := testApp.getHoldingsSlice()
	if len(slice) != 3 {
		t.Error("wrong number of rows returned")
	}
}
