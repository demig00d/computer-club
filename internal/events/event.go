package events

import (
	"errors"
	"fmt"
	"github.com/demig00d/computer-club/internal/client"
	"github.com/demig00d/computer-club/pkg/time24"
	"strconv"
	"strings"
)

var (
	ErrUnknownEventId  = errors.New("unknown event id")
	ErrNotEnoughFields = errors.New("not enough event's fields")
)

const (
	ClientCameIn = iota + 1
	ClientSatDownAtTheTable
	ClientIsWaiting
	ClientHasGone
	ClientHasGoneGen = iota + 7
	ClientSatDownAtTheTableGen
	ClientError
)

type Event struct {
	Time     time24.Time
	Id       int
	Client   client.Client
	TableId  int
	ErrorMsg string
}

func (e Event) String() string {
	switch e.Id {
	case ClientSatDownAtTheTable, ClientSatDownAtTheTableGen:
		return fmt.Sprintf("%s %d %s %d", e.Time.String(), e.Id, e.Client.Name, e.TableId)
	case ClientError:
		return fmt.Sprintf("%s %d %s", e.Time.String(), e.Id, e.ErrorMsg)
	default:
		return fmt.Sprintf("%s %d %s", e.Time.String(), e.Id, e.Client.Name)
	}
}

func Parse(s string) (Event, error) {
	fields := strings.Fields(s)
	if len(fields) < 3 {
		return Event{}, ErrNotEnoughFields
	}

	time, err := time24.Parse(fields[0])
	if err != nil {
		return Event{}, err
	}

	eventId, err := strconv.Atoi(fields[1])
	if err != nil {
		return Event{}, err
	}

	// ToDo возможно стоить перенести
	if eventId < 1 || eventId > 4 {
		return Event{}, ErrUnknownEventId
	}

	client, err := client.Parse(fields[2])

	if err != nil {
		return Event{}, err
	}

	var tableId int

	if len(fields) == 4 && (eventId == ClientSatDownAtTheTable || eventId == ClientSatDownAtTheTableGen) {
		tableId, err = strconv.Atoi(fields[3])
		if err != nil {
			return Event{}, err
		}

	}

	return Event{
		Time:    time,
		Id:      eventId,
		Client:  client,
		TableId: tableId,
	}, nil

}
