package ui

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func (app *Config) GetTickerText() (*canvas.Text, *canvas.Text) {
	var s Stock
	var currentPrice, symbol *canvas.Text

	p, err := s.GetPrices()
	if err != nil {
		grey := color.NRGBA{R: 155, G: 155, B: 155, A: 255}
		str := "Symbol: " + s.Symbol
		symbol = canvas.NewText(str, grey)
		currentPrice = canvas.NewText("Current Price: Network Unreachable", grey)
	}

	currentPriceTxt := fmt.Sprintf("Current Price: $%.4f", p.Current)
	currentPrice = canvas.NewText(currentPriceTxt, nil)

	symbolTxt := "Ticker: " + s.Symbol
	symbol = canvas.NewText(symbolTxt, nil)

	symbol.Alignment = fyne.TextAlignLeading
	currentPrice.Alignment = fyne.TextAlignCenter

	return symbol, currentPrice
}
