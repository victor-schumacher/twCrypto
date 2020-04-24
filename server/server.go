package server

type ReceivedData struct {
	Data struct {
		Base      string `json:"base"`
		Currency  string `json:"currency"`
		Amount    string `json:"amount"`
	} `json:"data"`
}

type (
	CryptoCurrency struct {
		Name      string
		BuyPrice  string
		SellPrice string
		Currency string
	}
	CryptoCurrencies[] CryptoCurrency
)
