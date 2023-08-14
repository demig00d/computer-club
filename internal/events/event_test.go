package events

import (
	"errors"
	"github.com/demig00d/computer-club/internal/client"
	"github.com/demig00d/computer-club/pkg/time24"
	"testing"
	"time"
)

var time00, _ = time24.Parse("00:00")

var testCases = []struct {
	input    string
	expected Event
}{
	{
		input: "09:00 1 client2",
		expected: Event{
			Time:   time00.Add(9 * time.Hour),
			Id:     1,
			Client: client.Client{Name: "client2"},
		},
	},
	{
		input: "11:00 2 client3 3",
		expected: Event{
			Time:    time00.Add(11 * time.Hour),
			Id:      2,
			Client:  client.Client{Name: "client3"},
			TableId: 3,
		},
	},
	{
		input: "19:00 4 client5",
		expected: Event{
			Time:   time00.Add(19 * time.Hour),
			Id:     4,
			Client: client.Client{Name: "client5"},
		},
	},
}

var InvalidEventsTestCases = []struct {
	eventString string
	expectedErr error
}{
	{
		eventString: "10:58 5 client3",
		expectedErr: ErrUnknownEventId,
	},
	{
		eventString: "10:58 2",
		expectedErr: ErrNotEnoughFields,
	},
}

func TestEvent(t *testing.T) {

	for _, test := range testCases {
		got, err := Parse(test.input)
		if err != nil {
			t.Errorf("unexpected error %s", err.Error())
		}

		if got != test.expected {
			t.Errorf("got:\n%v\n\nexpected:\n%v", got, test.expected)
		}
	}

	for _, test := range InvalidEventsTestCases {
		_, err := Parse(test.eventString)

		if !errors.Is(err, test.expectedErr) {
			t.Errorf("got: %s, expected: %s", err, test.expectedErr)
		}
	}

}
