package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
)

func deleteTimeLine(api *anaconda.TwitterApi, v url.Values) {
	tweets, err := api.GetUserTimeline(v)
	if err != nil {
		panic(err)
	}
	for _, tweet := range tweets {
		api.DeleteTweet(tweet.Id, true)
	}
}

func fuck(w http.ResponseWriter, r *http.Request) {
	const consumer_key = os.Getenv("CONSUMER_KEY")
	const consumer_secret = os.Getenv("CONSUMER_SECRET")
	const accsess_token = os.Getenv("ACCESS_TOKEN")
	const accsess_token_secret = os.Getenv("ACCESS_TOKEN_SECRET")
	const screen_name = os.Getenv("SCREEN_NAME")
	anaconda.SetConsumerKey(consumer_key)
	anaconda.SetConsumerSecret(consumer_secret)
	api := anaconda.NewTwitterApi(accsess_token, accsess_token_secret)
	v := url.Values{}
	v.Set("screen_name", screen_name)

	deleteTimeLine(api, v)
	fmt.Fprintf(w, "きえる")
}

func main() {
	http.HandleFunc("/", fuck)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}