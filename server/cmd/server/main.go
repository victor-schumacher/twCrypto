package main

import (
	"github.com/SchumacherVictor/twCrypto/server/crypto"
	"github.com/SchumacherVictor/twCrypto/server/twitter"
)

func main() {
	_ = "\"BTC, XLM, XMR, ETH, ZEC\""

	twitter.PostTweet(twitter.FormatTweet(crypto.GetCryptoData("ETH", "BRL")))






}
