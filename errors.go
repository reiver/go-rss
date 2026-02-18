package rss

import (
	"codeberg.org/reiver/go-erorr"
)

const (
	ErrNilReceiver       = erorr.Error("nil receiver")
	ErrPubDateUnparsable = erorr.Error("RSS pubDate unparsable")
)

const (
	errEmptyHTTPResponseBody   = erorr.Error("empty HTTP response body")
	errHTTPResponseStatusNotOK = erorr.Error("HTTP response status not OK")
	errNilHTTResponse          = erorr.Error("nil HTTP response")
)
