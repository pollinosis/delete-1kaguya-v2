package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

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
	anaconda.SetConsumerKey(CONSUMER_KEY)
	anaconda.SetConsumerSecret(CONSUMER_SECRET)
	api := anaconda.NewTwitterApi(ACCESS_TOKEN, ACCESS_TOKEN_SECRET)
	v := url.Values{}
	v.Set("screen_name", SCREEN_NAME)

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
