package main

import (
	"testing"
)

func TestApp_getToolBar(t *testing.T) {
	tb := testApp.getToolbar()

	if len(tb.Items) == 2 { // change to != for testing with real values
		t.Error("Toolbar construction failed")
	}
}
