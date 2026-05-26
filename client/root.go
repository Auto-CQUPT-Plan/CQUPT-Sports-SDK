package client

import (
	"net/http"

	persistentcookiejar "github.com/cdwiegand/persistent-cookiejar"
)

type Client struct {
	httpClient *http.Client
	cookieJar  *persistentcookiejar.Jar
}

func NewClient(path string) (*Client, error) {
	jar, err := persistentcookiejar.New(&persistentcookiejar.Options{Filename: path})
	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Jar:           jar,
		CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse },
	}

	s := &Client{httpClient: client, cookieJar: jar}

	return s, nil
}

func (r *Client) ShutdownClient() error {
	return r.cookieJar.Save()
}
