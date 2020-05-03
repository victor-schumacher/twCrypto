package crypto

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/SchumacherVictor/twCrypto/server"
)

func GetCrypto(cc string, c string) server.CryptoCurrency {
	var crypto server.CryptoCurrency

	crypto.Name = strings.ToUpper(cc)
	crypto.Currency = strings.ToUpper(c)
	crypto.SellPrice = GetPrice(cc, c, server.Operation.Sell)
	crypto.BuyPrice = GetPrice(cc, c, server.Operation.Buy)
	crypto.SpotPrice = GetPrice(cc, c, server.Operation.Spot)

	return crypto
}

func GetPrice(cc string, c string, op string) string {
	resp, err := http.Get(BuildURL(cc, c, op))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var rd server.ReceivedData
	err = json.Unmarshal(body, &rd)
	if err != nil {
		log.Fatalln(err)
	}

	return rd.Data.Amount

}

func BuildURL(cc string, c string, op string) string {
	base := "https://api.coinbase.com/v2/prices/"
	return base + strings.ToUpper(cc) + "-" + strings.ToUpper(c) + "/" + strings.ToLower(op)

}
