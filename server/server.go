package server

type (
	CryptoCurrency struct {
		SpotPrice      string
		BuyPrice  string
		SellPrice string
		Name      string
		Currency  string
	}
	CryptoCurrencies []CryptoCurrency

	ReceivedData struct {
		Data struct {
			Base     string `json:"base"`
			Currency string `json:"currency"`
			Amount   string `json:"amount"`
		} `json:"data"`
	}
	Op struct {
		Sell string
		Buy  string
		Spot string
	}
)

var Operation = Op{
	Buy:  "buy",
	Spot: "spot",
	Sell: "sell",
}
