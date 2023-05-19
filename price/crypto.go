package price

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const baseURL = "https://api.coingecko.com/api/v3/simple/price"

var coingeckoClient = http.Client{
	Timeout: 10 * time.Second,
}

type Price struct {
	Usd float64 `json:"usd"`
}

func Crypto(currencies ...string) map[string]Price {
	req := createRequest(currencies...)

	resp, _ := coingeckoClient.Do(req)

	var prices map[string]Price

	d := json.NewDecoder(resp.Body)
	d.Decode(&prices)

	return prices
}

// USD-TRY
func TRY() float64 {
	req := createRequest("tether")
	q := req.URL.Query()
	q.Set("vs_currencies", "try")
	q.Set("precision", "2")
	req.URL.RawQuery = q.Encode()

	resp, _ := coingeckoClient.Do(req)
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	p, _ := strconv.ParseFloat(string(body), 64)

	return p
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
