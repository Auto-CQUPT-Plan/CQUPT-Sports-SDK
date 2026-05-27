package CQUPT_Sports_SDK

import "github.com/Auto-CQUPT-Plan/CQUPT-Sports-SDK/client"

type Options func(s *Settings)

type Settings struct {
	Path string
}

type SDK struct {
	client *client.Client
}

func WithCookiePath(path string) Options {
	return func(s *Settings) {
		s.Path = path
	}
}

func NewSDK(opts ...Options) (*SDK, error) {
	settings := &Settings{
		Path: "./CookieJar.json",
	}

	for _, opt := range opts {
		opt(settings)
	}

	c, err := client.NewClient(settings.Path)
	if err != nil {
		return nil, err
	}

	s := &SDK{
		client: c,
	}

	return s, nil
}
