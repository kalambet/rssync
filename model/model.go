package model

import (
    "encoding/xml"
)

type Feed struct {
    XMLName xml.Name `xml:"rss"`
    Channel Channel `xml:"channel"`
}

type Channel struct {
    Title string `xml:"title"`
    Description string `xml:"description"`
    Link string `xml:"link"`
    Copyright string `xml:"copyright"`
    Items []Item `xml:"item"`

}

type Item struct {
    Title string `xml:"title"`
    Description string `xml:"description"`
    Link string `xml:"link"`
    GUID string `xml:"guid"`
    PubDate string `xml:"pubDate"`
    Watched bool
}