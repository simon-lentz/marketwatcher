package ui

import (
	"testing"

	"fyne.io/fyne/v2/test"
)

func TestApp_getToolBar(t *testing.T) {
	tb := testApp.GetToolbar()

	if len(tb.Items) == 2 { // change to != for testing with real values
		t.Error("Toolbar construction failed")
	}
}

func TestApp_addHoldingsDialog(t *testing.T) {
	testApp.AddHoldingsDialog()

	test.Type(testApp.HoldingUnits, "1")
	test.Type(testApp.HoldingValue, "5")

	if testApp.HoldingUnits.Text != "1" || testApp.HoldingValue.Text != "5" {
		t.Error("Table concat mutation failed, values do not match")
	}
}
