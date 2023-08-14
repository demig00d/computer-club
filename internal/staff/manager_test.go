package staff

import (
	"github.com/demig00d/computer-club/internal/computerclub"
	"github.com/demig00d/computer-club/internal/events"
	"github.com/demig00d/computer-club/internal/mockdata"
	"github.com/demig00d/computer-club/internal/workhours"
	"testing"
)

func MockManager() Manager {
	wh := workhours.NewWorkhours(mockdata.MockConfig())
	return NewManager(wh, MockEmployee())
}

func MockEmployee() Employee {
	mockComputerClub := computerclub.NewComputerClub(mockdata.MockConfig())
	return NewEmployee(mockComputerClub)
}

var (
	NotOpenAt2          = events.NotOpenYet(mockdata.Time02)
	ClientUnknownTime11 = events.ClientUnknown(mockdata.Time11)
)

var testCases = []struct {
	input    events.Event
	expected *events.Event
}{
	{
		input: events.Event{
			Time: mockdata.Time02,
			Id:   events.ClientCameIn,
		},
		expected: &NotOpenAt2,
	},
	{
		input: events.Event{
			Time: mockdata.Time11,
			Id:   events.ClientSatDownAtTheTable,
		},
		expected: nil,
	},

	{
		input: events.Event{
			Time: mockdata.Time11,
			Id:   events.ClientHasGone,
		},
		expected: &ClientUnknownTime11,
	},
}

func TestManage_HandleEvent(t *testing.T) {
	manager := MockManager()

	for _, test := range testCases {
		got := manager.HandleEvent(test.input)
		if !Equal(got, test.expected) {
			t.Errorf("got %v, but expected %v", got, test.expected)
		}
	}
}

func Equal(event1, event2 *events.Event) bool {
	if event1 == nil && event2 == nil {
		return true
	}
	if event1 == nil || event2 == nil {
		return false
	}

	return event1.Id == event2.Id &&
		event1.Client == event2.Client &&
		event1.TableId == event2.TableId &&
		event1.Time == event2.Time &&
		event1.ErrorMsg == event2.ErrorMsg

}
