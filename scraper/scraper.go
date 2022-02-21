package scraper

import (
	"fmt"
	"context"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

type Out struct {
	Currency string
	Price float64
}

func Scrape(ctx context.Context, url string) (Out, error) {
	return scrapeAmazon(ctx, url)
}

func scrapeAmazon(ctx context.Context, url string) (Out, error) {
	out := Out{}
	res, err := http.Get(url)
	if err != nil {
		return out, fmt.Errorf("Could not get %s: %w ", url, err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return out, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return out, fmt.Errorf("Could not load html document: %w", err)
	}
	priceStr, exists := doc.Find("input#twister-plus-price-data-price").First().Attr("value")
	if !exists {
		return out, fmt.Errorf("Could not find price tag with value attr")
	}
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return out, fmt.Errorf("Could not convert price %s to float64: %w", priceStr)
	}
	out.Price = price
	currency, exists := doc.Find("input#twister-plus-price-data-price-unit").First().Attr("value")
	if !exists {
		return out, fmt.Errorf("Could not find currency tag with value attr")
	}
	out.Currency = currency
	return out, nil
}

