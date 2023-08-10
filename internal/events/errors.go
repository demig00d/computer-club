package events

import (
	"github.com/demig00d/computer-club/internal/client"
	"github.com/demig00d/computer-club/pkg/time24"
)

func ErrEvent(time time24.Time, msg string) Event {
	return Event{
		Time:     time,
		Id:       13,
		Client:   client.Client{},
		TableId:  0,
		ErrorMsg: msg,
	}

}

func NotOpenYet(time time24.Time) Event {
	return ErrEvent(time, "NotOpenYet")
}

func YouShallNotPass(time time24.Time) Event {
	return ErrEvent(time, "YouShallNotPass")
}

func ICanWaitNoLonger(time time24.Time) Event {
	return ErrEvent(time, "ICanWaitNoLonger!")
}

func PlaceIsBusy(time time24.Time) Event {
	return ErrEvent(time, "PlaceIsBusy")
}

func ClientUnknown(time time24.Time) Event {
	return ErrEvent(time, "ClientUnknown")
}
