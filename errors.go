package rss

import (
	"codeberg.org/reiver/go-erorr"
)

const (
	ErrNilReceiver       = erorr.Error("nil receiver")
	ErrPubDateUnparsable = erorr.Error("RSS pubDate unparsable")
)
