package main

import (
	"fmt"
	"github.com/SchumacherVictor/twCrypto/server/crypto"
	"github.com/SchumacherVictor/twCrypto/server/twitter"
	"github.com/roylee0704/gron"
	"log"
	"net/http"
	"time"
)

func main() {

	var cc = [5] string {"BTC", "ETH", "XRP", "LTC", "BCH" }
	i := 0
	c := gron.New()
	c.AddFunc(gron.Every(1*time.Second), func() {
		//twitter.PostTweet(twitter.FormatTweet(crypto.GetCryptoData(cc[i], "USD")))
		fmt.Println(twitter.FormatTweet(crypto.GetCryptoData(cc[i], "USD")))
		i++
		if i == 5 {i=0}
	})
	c.Start()


	log.Fatal(http.ListenAndServe(":8080", nil))

}
