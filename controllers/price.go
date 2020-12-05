package controllers

import (
	"bytes"
	"github.com/antchfx/htmlquery"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"net/http"
)

// company is a structure that contains the company's stock ticker from the client's HTTP request
type company struct {
	Ticker string `json:"ticker" form:"ticker" query:"ticker"`
}
type price struct {
	Price string `json:"price" form:"price" query:"price"`
}

// GrabPrice - handler method for binding JSON body and scraping for stock price
func GrabPrice(c echo.Context) (err error) {
	// Read the body content
	var bodyBytes []byte
	if c.Request().Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request().Body)
	}

	c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	u := new(company)
	er := c.Bind(u) // bind the structure with the context body
	// on no panic!
	if er != nil {
		panic(er)
	}

	// Company ticker
	ticker := u.Ticker

	// yahoo finance base URL
	baseURL := "https://finance.yahoo.com/quote/"

	// price Xpath
	pricePath := "//*[@id=\"quote-header-info\"]"

	// load HTML document by binding base url and passed in ticker
	doc, err := htmlquery.LoadURL(baseURL + ticker)
	// uh oh :( freak out!
	if err != nil {
		panic(err)
	}

	// HTML Node
	context := htmlquery.FindOne(doc, pricePath)

	// from the Node get inner text
	var p price
	p.Price = htmlquery.InnerText(context)
	return c.JSON(http.StatusOK, p)
}
