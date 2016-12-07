package main

import (
	"encoding/xml"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://soap4.me/rss/soap-429/quality-2/translate-999/0503d3d80fea6fa3011ea3664f7e1719/")

	if err != nil {
		log.Fatal("Error:" + err.Error())
		return
	}

	defer resp.Body.Close()
	xmlDecoder := xml.NewDecoder(resp.Body)

	log.Print(body)
}
