package converter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	apiKeyEnv = "COINMARKETCAP_API_KEY"
	apiUrlEnv = "COINMARKETCAP_API_URL"
)

type apiResponse struct {
	Data struct {
		ID     int    `json:"id"`
		Symbol string `json:"symbol"`
		Name   string `json:"name"`
		Quote  map[string]struct {
			Price float64 `json:"price"`
		} `json:"quote"`
	} `json:"data"`
}

type apiErrorResponse struct {
	Status struct {
		ErrorMessage string `json:"error_message"`
	} `json:"status"`
}

func GetConversionRate(fromCurrency, toCurrency string) (float64, error) {
	apiKey := os.Getenv(apiKeyEnv)
	if apiKey == "" {
		return 0, fmt.Errorf("environment variable %s not set", apiKeyEnv)
	}

	apiUrl := os.Getenv(apiUrlEnv)
	if apiUrl == "" {
		return 0, fmt.Errorf("environment variable %s not set", apiUrlEnv)
	}

	url := fmt.Sprintf("%s?amount=1&symbol=%s&convert=%s", apiUrl, fromCurrency, toCurrency)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("X-CMC_PRO_API_KEY", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	if resp.StatusCode != http.StatusOK {
		var apiErr apiErrorResponse
		if err := json.Unmarshal(body, &apiErr); err != nil {
			return 0, fmt.Errorf("failed to get data: %s", resp.Status)
		}
		return 0, fmt.Errorf("failed to get data: %s, %s", resp.Status, apiErr.Status.ErrorMessage)
	}

	var apiResp apiResponse
	if err := json.Unmarshal(body, &apiResp); err != nil {
		return 0, fmt.Errorf("failed to parse JSON response: %v", err)
	}

	rate, ok := apiResp.Data.Quote[toCurrency]
	if !ok {
		return 0, fmt.Errorf("conversion rate not found for %s to %s", fromCurrency, toCurrency)
	}

	return rate.Price, nil
}
