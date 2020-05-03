package twitter

import (
	"github.com/SchumacherVictor/twCrypto/server"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func FormatTweet(c server.CryptoCurrency) string {
	t := "Cryptocurrency ➔ " + c.Name + "\n" +
		"Spot ➔ " + c.SpotPrice + "$\n" +
		"Sell ➔ " + c.SellPrice + "$\n" +
		"Buy ➔ " + c.BuyPrice + "$\n" +
		"#crypocurrencies 😆\n#" + c.Name
	return t
}

func PostTweet(t string) {
	config := oauth1.NewConfig(os.Getenv("consumerKey"), os.Getenv("consumerSecret"))
	token := oauth1.NewToken(os.Getenv("accessToken"), os.Getenv("accessTokenSecret"))
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	client.Statuses.Update(t, nil)
}
