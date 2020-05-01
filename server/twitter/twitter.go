package twitter

import (
	"os"

	"github.com/SchumacherVictor/twCrypto/server"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func FormatTweet(c server.CryptoCurrency) string {
	finalTweet := "CryptoC    |        Sell         |         Buy         |\n" +
		"   " + c.Name + "          "    + c.SellPrice + "$        "    + c.BuyPrice +"$" +
		"\n #testing"
	return finalTweet
}

func PostTweet(t string) {
	config := oauth1.NewConfig(os.Getenv("consumerKey"), os.Getenv("consumerSecret"))
	token := oauth1.NewToken(os.Getenv("accessToken"), os.Getenv("accessTokenSecret"))
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	client.Statuses.Update(t, nil)
}
