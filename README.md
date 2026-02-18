# go-rss

Package **rss** implements RSS encoders and decoders, for the Go programming language.

Package **rss** is meant to be used with the Go built-in `"encoding/xml"` package.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-rss

[![GoDoc](https://godoc.org/github.com/reiver/go-rss?status.svg)](https://godoc.org/github.com/reiver/go-rss)

## Examples

To unmarshal RSS data to `rss.RSS2` in Go, you can do something similar to the following:

```golang
import (
	"encoding/xml"

	"github.com/reiver/go-rss"
}

// ...

var rss2 rss.RSS2

err := xml.Unmarshal(p, &rss2)
```

To make an HTTP request to a URL for RSS and then unmarshal it, you can do something similar to the following:

```golang
import (
	"github.com/reiver/go-rss"
}

// ...

url := "https://mastodon.social/tags/fedidev.rss"

var rss2 rss.RSS2

err := rss.HTTPGetAndUnmarshal(url, &rss2)
```

## Installation

To install package **rss** do the following:
```
GOPROXY=direct go get github.com/reiver/go-rss
```

## Author

Package **rss** was written by [Charles Iliya Krempeaux](http://reiver.link)
