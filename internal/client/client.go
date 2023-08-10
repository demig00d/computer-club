package client

import (
	"errors"
	"regexp"
)

var (
	ErrInvalidCharactersInClientName = errors.New("invalid characters in the client's name, should be a-zA-Z0-9_-")
)

type Client struct {
	Name string
}

func Parse(s string) (Client, error) {
	m, err := regexp.MatchString("^[a-zA-Z0-9_\\-]+$", s)
	if err != nil {
		return Client{}, err
	}

	if m {
		return Client{
			Name: s,
		}, nil
	}

	return Client{}, ErrInvalidCharactersInClientName
}
