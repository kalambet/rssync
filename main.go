package main

import (
	"log"

	"github.com/kalambet/rssync/reader"
)

func main() {
	data, err := reader.ReadRSSFeed("http://soap4.me/rss/soap-429/quality-2/translate-999/0503d3d80fea6fa3011ea3664f7e1719/")

	if err != nil {
		log.Fatal("Error:" + err.Error())
	}

	for _, item := range data.Channel.Items {
		//url, _ := url.Parse(item.Link)
		log.Printf("PubDate: %#v", item.Title)
	}
}
