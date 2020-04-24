package crypto

import (
	"cryptoTwitter/server"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func GetCryptoData(cc string, c string) server.CryptoCurrency {
	op := "sell"
	resp, err := http.Get(BuildUrl(cc,c,op))
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

	var crypto server.CryptoCurrency
	crypto.Name = rd.Data.Base
	crypto.Currency = rd.Data.Currency
	crypto.SellPrice = rd.Data.Amount

	op2 := "buy"
	resp2, err := http.Get(BuildUrl(cc,c,op2))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp2.Body.Close()

	body2, err := ioutil.ReadAll(resp2.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var rd2 server.ReceivedData
	err = json.Unmarshal(body2, &rd2)
	if err != nil {
		log.Fatalln(err)
	}

	crypto.BuyPrice = rd2.Data.Amount
	return crypto

}

func BuildUrl(cc string, c string, op string) string {
	//read about const in golang and move this strings to one place
	base := "https://api.coinbase.com/v2/prices/"
	return base + strings.ToUpper(cc) + "-" + strings.ToUpper(c) + "/" + strings.ToLower(op)

}
