package rss

import (
	"encoding/xml"
	"fmt"
	"time"
)

// RSS2 represents a full RSS 2.0 XML document.
type RSS2 struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

type Channel struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Items       []Item `xml:"item"`
}

type Item struct {
	Author      string      `xml:"author"`
	Categories  []string    `xml:"category"`
	Description Description `xml:"description"`
	GUID        GUID        `xml:"guid"`
	Link        string      `xml:"link"`
	PubDate     PubDate     `xml:"pubDate"`
	Title       string      `xml:"title"`
}

type Description struct {
	Value string `xml:",chardata"`
}

func (receiver *Description) HTML() string {
	if nil == receiver {
		panic(ErrNilReceiver)
	}

	//@TODO
	return receiver.Value
}

type GUID struct {
	IsPermaLink bool   `xml:"isPermaLink,attr"`
	Value       string `xml:",chardata"`
}

type PubDate struct {
	Value string `xml:",chardata"`
}

var timeFormats = []string{
	time.RFC1123Z,
	time.RFC1123,
	"Mon, 02 Jan 2006 15:04:05 -0700",
	"Mon, 02 Jan 2006 15:04:05 MST",
}

func (receiver PubDate) Parse() (time.Time, error) {

	for _, format := range timeFormats {
		t, err := time.Parse(format, receiver.Value)
		if nil == err {
			return t, nil
		}
	}

	return time.Time{}, fmt.Errorf("unable to parse RSS pubDate %q: %w", receiver.Value, ErrPubDateUnparsable)
}
