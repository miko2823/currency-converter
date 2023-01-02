package converter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/miko2823/currency-converter.git/domain/converter/models"
)

func (r converterPersistence) GetLatestRates(base string, symbols string, amount int) (*models.LatestRates, error) {

	url := fmt.Sprintf("https://api.apilayer.com/exchangerates_data/latest?symbols=%s&base=%s", symbols, base)
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("apikey", r.config.CONVERTER_API_KEY)
	res, err := client.Do(req)
	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var converterResponse map[string]interface{}
	latestRate := models.LatestRates{}

	if err := json.Unmarshal([]byte(string(body)), &converterResponse); err != nil {
		return nil, err
	}

	latestRate.Base = converterResponse["base"].(string)
	latestRate.Amount = amount

	for key, val := range converterResponse["rates"].(map[string]interface{}) {
		exchangeAmount := 1.0 / val.(float64) * float64(amount)
		latestRate.Rates = append(latestRate.Rates, models.Rate{string(key), exchangeAmount})
	}

	return &latestRate, nil
}
