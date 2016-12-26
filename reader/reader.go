package reader

import (
	"encoding/xml"
	"net/http"

	"github.com/kalambet/rssync/model"
)

// ReadRssFeed imports rss feed content and created intenal data model instnace
func ReadRssFeed(url string) (data *model.Feed, err error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	//defer file.Close()
	decoder := xml.NewDecoder(resp.Body)

	data = &model.Feed{}
	err = decoder.Decode(data)

	if err != nil {
		return nil, err
	}

	return data, nil
}
