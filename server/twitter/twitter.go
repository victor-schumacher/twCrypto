package twitter

import (
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func FormatTweet(cryptocurrency string, sellPrice string, buyPrice string) string {
	t := "Cryptocurrency ➔ " + cryptocurrency + "\n" +
		 "Sell price ➔ "+sellPrice+ "$\n" +
		 "Buy price ➔ "+buyPrice+"$\n"+
		 "#crypocurrencies\n#"+cryptocurrency
	return t
}

func PostTweet(t string) {
	config := oauth1.NewConfig(os.Getenv("consumerKey"), os.Getenv("consumerSecret"))
	token := oauth1.NewToken(os.Getenv("accessToken"), os.Getenv("accessTokenSecret"))
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	client.Statuses.Update(t, nil)
}
