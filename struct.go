package main

import (
	"encoding/xml"
)

type Response struct {
	XMLName xml.Name `xml:"feed"`
	Text    string   `xml:",chardata"`
	Xmlns   string   `xml:"xmlns,attr"`
	Media   string   `xml:"media,attr"`
	Lang    string   `xml:"lang,attr"`
	ID      string   `xml:"id"`
	Link    []struct {
		Text string `xml:",chardata"`
		Type string `xml:"type,attr"`
		Rel  string `xml:"rel,attr"`
		Href string `xml:"href,attr"`
	} `xml:"link"`
	Title   string `xml:"title"`
	Updated string `xml:"updated"`
	Entry   []struct {
		Text      string `xml:",chardata"`
		ID        string `xml:"id"`
		Published string `xml:"published"`
		Updated   string `xml:"updated"`
		Link      struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
			Rel  string `xml:"rel,attr"`
			Href string `xml:"href,attr"`
		} `xml:"link"`
		Title struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
		} `xml:"title"`
		Author struct {
			Text string `xml:",chardata"`
			Name string `xml:"name"`
			URI  string `xml:"uri"`
		} `xml:"author"`
		Thumbnail struct {
			Text   string `xml:",chardata"`
			Height string `xml:"height,attr"`
			Width  string `xml:"width,attr"`
			URL    string `xml:"url,attr"`
		} `xml:"thumbnail"`
		Content struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
		} `xml:"content"`
	} `xml:"entry"`
}

type Metadata struct {
	TweetDate    string
	CurrentTopId int
	Templates    []Template
}

type Template struct {
	ID    int
	Event string
	Title string
	URL   string
	Date  string
}
