package staff

import (
	"fmt"
	"github.com/demig00d/computer-club/config"
	"github.com/demig00d/computer-club/internal/computerclub"
	"github.com/demig00d/computer-club/internal/events"
)

type Manager struct {
	cfg         config.Config
	subordinate Employee
}

func NewManager(cfg config.Config) Manager {
	club := computerclub.NewComputerClub(cfg)
	return Manager{
		cfg:         cfg,
		subordinate: NewEmployee(club),
	}
}

func (m Manager) OpenClub() {
	fmt.Println(m.cfg.OpeningTime())
}

func (m Manager) CloseClub() {

	es := m.subordinate.KickOutClients()
	for _, e := range es {
		fmt.Println(e)
		m.subordinate.EscortClientOut(e.Client, e.Time)
	}

	endTime := m.cfg.ClosingTime()
	fmt.Println(endTime)

	// для всех столов выводим выручку и общее время
	report := m.subordinate.FormTablesReport()
	fmt.Print(report)
}

func (m Manager) HandleEvent(event events.Event) *events.Event {
	var newEvent *events.Event

	switch event.Id {
	case events.ClientCameIn:
		newEvent = m.subordinate.MeetClient(event.Client, event.Time)
	case events.ClientSatDownAtTheTable, events.ClientSatDownAtTheTableGen:
		newEvent = m.subordinate.SeatClientAt(event.Client, event.Time, event.TableId)
	case events.ClientIsWaiting:
		newEvent = m.subordinate.AddClientToQueue(event.Client, event.Time)
	case events.ClientHasGone, events.ClientHasGoneGen:
		newEvent = m.subordinate.EscortClientOut(event.Client, event.Time)
	}
	return newEvent

}
