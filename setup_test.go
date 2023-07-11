package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"testing"

	"fyne.io/fyne/v2/test"
	"github.com/simon-lentz/marketwatcher/repo"
)

var testApp Config

/*
Since testing would require an active network
and the target site being up, instead we reroute
references to http.Client a local value so that
the logic can be tested in isolation from the
web content.

In this case, I just took a demo API response
matching the desired behavior for tickers and
used that to configure the testing environment.
*/
func TestMain(m *testing.M) {
	// setup testing environment
	a := test.NewApp()
	testApp.App = a
	testApp.MainWindow = a.NewWindow("")
	testApp.HTTPClient = client
	testApp.DB = repo.NewTestRepo()
	os.Exit(m.Run()) // tests run after this
}

var jsonToReturn = `
{
    "Global Quote": {
        "01. symbol": "IBM",
        "02. open": "131.7600",
        "03. high": "133.0500",
        "04. low": "131.6950",
        "05. price": "132.9000",
        "06. volume": "2369425",
        "07. latest trading day": "2023-07-10",
        "08. previous close": "132.0800",
        "09. change": "0.8200",
        "10. change percent": "0.6208%"
    }
}
`

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}

var client = NewTestClient(func(req *http.Request) *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString(jsonToReturn)),
		Header:     make(http.Header),
	}
})
