package workhours

import (
	"github.com/demig00d/computer-club/config"
	"github.com/demig00d/computer-club/pkg/time24"
)

type Workhours struct {
	openingTime time24.Time
	closingTime time24.Time
}

func (w Workhours) OpeningTime() time24.Time {
	return w.openingTime
}

func (w Workhours) ClosingTime() time24.Time {
	return w.closingTime
}

func NewWorkhours(cfg config.Config) Workhours {
	return Workhours{
		openingTime: cfg.OpeningTime(),
		closingTime: cfg.ClosingTime(),
	}
}
