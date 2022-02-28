package poster

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/amirkhaki/cmuxbtr/config"
	"github.com/amirkhaki/cmuxbtr/scraper"
	"io"
	"math"
	"net/http"
	"strconv"
	"strings"
)

func addTax(price, taxPercent float64) float64 {
	return ((price / 100) * taxPercent) + price
}

func Post(ctx context.Context, id int, out scraper.Out, tax, wage float64) error {
	cfg := config.Cfg
	if strings.ToLower(out.Currency) != "aed" {
		return fmt.Errorf("Unsupported currency, only AED is valid!")
	}
	salePrice := out.SalePrice * cfg.AEDPrice
	salePrice = addTax(salePrice, tax) + wage
	regularPrice := out.RegularPrice * cfg.AEDPrice
	regularPrice = addTax(regularPrice, tax) + wage
	priceDto := struct {
		RegularPrice string `json:"regular_price"`
		SalePrice    string `json:"sale_price"`
	}{
		RegularPrice: strconv.Itoa(int(math.Round(regularPrice/1000)*1000)),
		SalePrice:    strconv.Itoa(int(math.Round(salePrice/1000)*1000)),
	}
	jsonByte, err := json.Marshal(priceDto)
	if err != nil {
		return fmt.Errorf("Error during marshaling json body: %w", err)
	}
	wc_up_prdct_endpoint := cfg.WPURL + "/wp-json/wc/v3/products/"
	wc_up_prdct_endpoint += fmt.Sprintf("%d", id)
	wc_up_prdct_endpoint += "?consumer_key=" + cfg.WPKey + "&consumer_secret=" + cfg.WPSecret
	req, err := http.NewRequest("POST", wc_up_prdct_endpoint, bytes.NewBuffer(jsonByte))
	if err != nil {
		return fmt.Errorf("Error during creating request: %w", err)
	}
	req.SetBasicAuth(cfg.WPKey, cfg.WPSecret)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Error during posting product: %w", err)
	}
	defer resp.Body.Close()
	response, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Error during reading response body: %w", err)
	}
	fmt.Println(string(response))
	fmt.Println(priceDto)

	return nil
}
