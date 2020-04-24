package twitter

import "github.com/dghubble/go-twitter/twitter"
import "github.com/dghubble/oauth1"



func postTweet(){
	config := oauth1.NewConfig("consumerKey", "consumerSecret")
	token := oauth1.NewToken("accessToken", "accessSecret")
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)


}

