package client

import (
	"errors"
	"testing"
)

var testTable = []struct {
	input    string
	expected Client
}{
	{
		input:    "client2",
		expected: Client{"client2"},
	},
	{
		input:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz-0123456789",
		expected: Client{"ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz-0123456789"},
	},
}

var InvalidNames = []struct {
	name        string
	expectedErr error
}{
	{
		name:        "Клиент5",
		expectedErr: ErrInvalidCharactersInClientName,
	},
	{
		name:        "client#",
		expectedErr: ErrInvalidCharactersInClientName,
	},
	{
		name:        ">",
		expectedErr: ErrInvalidCharactersInClientName,
	},
}

func TestEvent(t *testing.T) {

	for _, test := range testTable {
		got, err := Parse(test.input)
		if err != nil {
			t.Errorf("unexpected error %s", err.Error())
		}

		if got != test.expected {
			t.Errorf("got:\n%v\n\nexpected:\n%v", got, test.expected)
		}
	}

	for _, test := range InvalidNames {
		_, err := Parse(test.name)

		if !errors.Is(err, test.expectedErr) {
			t.Errorf("got: %s, expected: %s", err, test.expectedErr)
		}
	}

}
