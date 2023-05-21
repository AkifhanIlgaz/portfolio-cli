package price

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

const baseURL = "https://api.coingecko.com/api/v3/simple/price"

var httpClient = http.Client{
	Timeout: 10 * time.Second,
}

type PriceData map[string]map[string]float64

func (p PriceData) GetPrice() float64 {
	var price float64
	/*
		{
			"sui":
				{
					"usd":1.1284195321
				}
		}
		Since our price API is making request for one currency, iterating through PriceData map gives us the USD price.
		We can achieve the same goal by taking currency name as argument
	*/
	for _, v := range p {
		price = v["usd"]
	}

	return price
}

func Crypto(currency string) PriceData {
	req := createRequest(currency)

	resp, _ := httpClient.Do(req)

	var price PriceData

	d := json.NewDecoder(resp.Body)
	d.Decode(&price)

	return price

}

// USD-TRY
func TRY() float64 {
	resp, _ := httpClient.Get("https://hasanadiguzel.com.tr/api/kurgetir")

	defer resp.Body.Close()

	type CurrencyInfo struct {
		Developer struct {
			Name    string `json:"name"`
			Website string `json:"website"`
			Email   string `json:"email"`
		} `json:"developer"`
		TCMBAnlikKurBilgileri []struct {
			Isim            string `json:"Isim"`
			CurrencyName    string `json:"CurrencyName"`
			ForexBuying     any    `json:"ForexBuying"`
			ForexSelling    any    `json:"ForexSelling"`
			BanknoteBuying  any    `json:"BanknoteBuying"`
			BanknoteSelling any    `json:"BanknoteSelling"`
			CrossRateUSD    any    `json:"CrossRateUSD"`
			CrossRateOther  any    `json:"CrossRateOther"`
		} `json:"TCMB_AnlikKurBilgileri"`
	}

	var info CurrencyInfo
	json.NewDecoder(resp.Body).Decode(&info)

	return (info.TCMBAnlikKurBilgileri[0].ForexSelling).(float64)
}

func createRequest(currency string) *http.Request {

	q := url.Values{}
	q.Set("ids", currency)
	q.Set("vs_currencies", "usd")
	q.Set("precision", "5")

	req, _ := http.NewRequest("GET", baseURL, nil)
	req.URL.RawQuery = q.Encode()

	return req
}
