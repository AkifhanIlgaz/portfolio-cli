package price

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const baseURL = "https://api.coingecko.com/api/v3/simple/price"

var httpClient = http.Client{
	Timeout: 10 * time.Second,
}

type PriceData map[string]Price

type Price struct {
	Usd float64
}

func Crypto(currencies ...string) PriceData {
	req := createRequest(currencies...)

	resp, _ := httpClient.Do(req)

	var prices PriceData

	d := json.NewDecoder(resp.Body)
	d.Decode(&prices)

	return prices

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

func createRequest(currencies ...string) *http.Request {

	q := url.Values{}
	q.Set("ids", strings.Join(currencies, ","))
	q.Set("vs_currencies", "usd")
	q.Set("precision", "10")

	req, _ := http.NewRequest("GET", baseURL, nil)
	req.URL.RawQuery = q.Encode()

	return req
}
