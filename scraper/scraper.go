package scraper

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/corpix/uarand"
	"net/http"
	"strconv"
	"strings"
)

type Out struct {
	Currency     string
	SalePrice    float64
	RegularPrice float64
}

func Scrape(ctx context.Context, url string) (Out, error) {
	return scrapeAmazon(ctx, url)
}

func scrapeAmazon(ctx context.Context, url string) (Out, error) {
	out := Out{}
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return out, fmt.Errorf("Could not create request: %w ", err)
	}
	req.Header.Set("User-Agent", uarand.GetRandom())
	res, err := client.Do(req)
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
	salePrice, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return out, fmt.Errorf("Could not convert price %s to float64: %w", priceStr, err)
	}
	out.SalePrice = salePrice
	currency, exists := doc.Find("input#twister-plus-price-data-price-unit").First().Attr("value")
	if !exists {
		return out, fmt.Errorf("Could not find currency tag with value attr")
	}
	out.Currency = currency
	listPriceSelector := "span.basisPrice span.a-price.a-text-price span.a-offscreen"
	listPriceText := doc.Find(listPriceSelector).First().Text()
	fmt.Println(listPriceText)
	if listPriceText == "" {
		out.RegularPrice = out.SalePrice
	} else {
		listPriceText = strings.Replace(listPriceText, "AED", "", -1)
		listPriceText = strings.Replace(listPriceText, ",", "", -1)
		listPrice, err := strconv.ParseFloat(listPriceText, 64)
		if err != nil {
			return out, fmt.Errorf("Could not convert list price %s to float64: %w", listPriceText, err)
		}
		out.RegularPrice = listPrice
	}
	return out, nil
}
