package crypto

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/SchumacherVictor/twCrypto/server"
)

func GetCryptoData(cc string, c string) server.CryptoCurrency {

	var crypto server.CryptoCurrency
	crypto.Name = strings.ToUpper(cc)
	crypto.Currency = strings.ToUpper(c)
	crypto.SellPrice = GetSellPrice(cc, c)
	crypto.BuyPrice = GetBuyPrice(cc, c)
	return crypto

}
func GetSellPrice(cc string, c string) string {

	resp, err := http.Get(BuildURL(cc, c, "sell"))
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

func GetBuyPrice(cc string, c string) string {

	resp, err := http.Get(BuildURL(cc, c, "buy"))
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
