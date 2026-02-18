package rss

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

// HTTPGet makes an HTTP GET request to a URL with the HTTP "Accept" request header set to "application/rss+xml" (the RSS media-type).
//
// Example usage:
//
//	bytes, err := rss.HTTPGet(url)
//
// See also: [HTTPGetAndUnmarshal]
func HTTPGet(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if nil != err {
		return nil, fmt.Errorf("failed to create new HTTP GET request (that will be to %q to get RSS): %w", url, err)
	}
	req.Header.Set("Accept", MediaType)

	httpClient := http.DefaultClient

	resp, err := httpClient.Do(req)
	if nil != err {
		return nil, fmt.Errorf("failed making HTTP GET request to %q to (try to) get RSS: %w", url, err)
	}
	if nil == resp {
		err = errNilHTTResponse
		return nil, fmt.Errorf("failed making HTTP GET request to %q to (try to) get RSS: %w", url, err)
	}

	if http.StatusOK != resp.StatusCode {
		err = errHTTPResponseStatusNotOK
		return nil, fmt.Errorf("failed to get \"200 OK\" HTTP response from %q (which should hae been RSS); actually got %d: %w", url, resp.StatusCode, err)
	}

	body, err := io.ReadAll(resp.Body)
	if nil != err {
		return nil, fmt.Errorf("failed to read HTTP response body from the URL %q that should have been RSS: %w", url, err)
	}
	if len(body) <= 0 {
		err = errEmptyHTTPResponseBody
		return nil, fmt.Errorf("failed to xml-unmarshal what should have been RSS from the URL %q: %w", url, err)
	}

	return body, nil
}

// HTTPGetAndMarshal makes an HTTP GET request to a URL with the HTTP "Accept" request header set to "application/rss+xml" (the RSS media-type),
// and then unmarshals the body of the resulting HTTP response into RSS.
//
// Example usage:
//
//	var rss2 rss.RSS2
//	err := rss.HTTPGetAndUnmarshal(url, &rss2)
//
// See also: [HTTPGet]
func HTTPGetAndUnmarshal(url string, dst any) error {
	body, err := HTTPGet(url)
	if nil != err {
		return err
	}

	if nil == dst {
		return nil
	}

	err = xml.Unmarshal(body, dst)
	if nil != err {
		return fmt.Errorf("failed to xml-unmarshal what should have been RSS from the URL %q into %T: %w", url, dst, err)
	}

	return nil
}
