package main

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/png"
	"io"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func (app *Config) pricesTab() *fyne.Container {
	img := app.getImageAsset()
	imgContainer := container.NewVBox(img)
	app.PriceImgContainer = imgContainer

	return imgContainer
}

func (app *Config) getImageAsset() *canvas.Image {

	apiURL := fmt.Sprintf("")
	var img *canvas.Image

	if err := app.downloadFile(apiURL, "asset.png"); err != nil {
		// use bundled image (static asset) ./fyne bundled filename.png >> filename.go
		img = canvas.NewImageFromResource(resourceHttp404Png)
	} else {
		img = canvas.NewImageFromFile("asset.png")
	}

	img.SetMinSize(fyne.Size{
		Width:  250,
		Height: 250,
	})

	img.FillMode = canvas.ImageFillOriginal

	return img
}

func (app *Config) downloadFile(URL, fileName string) error {
	// get resp bytes from req
	resp, err := app.HTTPClient.Get(URL)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New("Unexpected status code: " + fmt.Sprint(resp.StatusCode))
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return err
	}

	out, err := os.Create(fmt.Sprintf("./%s", fileName))

	if err = png.Encode(out, img); err != nil {
		return err
	}

	return nil
}
