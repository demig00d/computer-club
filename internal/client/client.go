package client

import (
	"errors"
	"regexp"
)

type Client struct {
	Name string
}

func Parse(s string) (Client, error) {
	m, err := regexp.MatchString("\\w+", s)
	if err != nil {
		return Client{}, err
	}

	if m {
		return Client{
			Name: s,
		}, nil
	}

	return Client{}, errors.New("incorrect name format, should be a-zA-Z0-9_-")
}
