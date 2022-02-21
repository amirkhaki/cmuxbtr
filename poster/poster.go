package poster

import (
	"github.com/amirkhaki/cmuxbtr/config"
	"github.com/amirkhaki/cmuxbtr/scraper"
	"context"
	"strings"
	"fmt"
	"net/http"
	"bytes"
)

func Post(ctx context.Context, cfg *config.Config, id int, out scraper.Out) error {
	if strings.ToLower(out.Currency) != "aed" {
		return fmt.Errorf("Unsupported currency, only AED is valid!")
	}
	price := out.Price * cfg.AEDPrice 
	jsonStr := fmt.Sprintf(`{"regular_price":%d}`, price)
	wc_up_prdct_endpoint := cfg.WPURL + "/wp-json/wc/v3/products/"
	wc_up_prdct_endpoint += fmt.Sprintf("%d", id)
	wc_up_prdct_endpoint += "?consumer_key=" +  cfg.WPKey + "&consumer_secret=" + cfg.WPSecret
	req, err := http.NewRequest("POST", wc_up_prdct_endpoint, bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		return fmt.Errorf("Error during creating request: %w", err)
	}
	req.SetBasicAuth(cfg.WPKey, cfg.WPSecret )
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		return fmt.Errorf("Error during posting product: %w", err)
	}
	return nil
}
