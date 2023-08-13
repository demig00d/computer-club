package staff

import (
	"errors"
	"fmt"
	"github.com/demig00d/computer-club/internal/client"
	"github.com/demig00d/computer-club/internal/computerclub"
	"github.com/demig00d/computer-club/internal/events"
	"github.com/demig00d/computer-club/internal/tables"
	"github.com/demig00d/computer-club/pkg/time24"
	"sort"
	"strings"
)

// Employee использует состояние и методы ComputerClub,
// чтобы исполнять действия присланные событиями (Event)
type Employee struct {
	club *computerclub.ComputerClub
}

func NewEmployee(club computerclub.ComputerClub) Employee {
	return Employee{&club}
}

func (e Employee) MeetClient(client client.Client, time time24.Time) *events.Event {
	if e.club.IsBeforeOpening(time) {
		event := events.NotOpenYet(time)
		return &event
	}

	if e.club.IsClientIn(client) {
		event := events.YouShallNotPass(time)
		return &event
	}
	return nil
}

func (e Employee) AddClientToQueue(client client.Client, time time24.Time) *events.Event {
	if e.club.AreThereAnyTablesAvailable() {
		event := events.ICanWaitNoLonger(time)
		return &event
	}

	if e.club.AreClientsInQueueMoreThanTables() {
		return &events.Event{
			Time:   time,
			Id:     11,
			Client: client,
		}
	}

	_ = e.club.ClientQueue.Push(client)

	return nil
}

// SeatClientAt Добавляет Client в Table
func (e Employee) SeatClientAt(client client.Client, time time24.Time, tableId int) *events.Event {

	if e.club.IsClientIn(client) {
		event := events.ClientUnknown(time)
		return &event
	}

	table := e.club.Tables.GetTable(tableId)
	if table == nil {
		return nil
	}

	if !table.IsEmpty() {
		event := events.PlaceIsBusy(time)
		return &event
	}

	table.SeatClient(client, time)
	e.club.ClientQueue.Remove(client)

	return nil
}

func (e Employee) EscortClientOut(client client.Client, time time24.Time) *events.Event {
	if !e.club.IsClientIn(client) {
		event := events.ClientUnknown(time)
		return &event
	}

	tableId, err := e.vacateTheTable(client, time)
	if err != nil {
		e.club.ClientQueue.Remove(client)
		return nil
	}

	waitingClient, err := e.club.ClientQueue.Poll()

	if err != nil {
		return nil
	}

	return &events.Event{
		Time:    time,
		Id:      12,
		Client:  waitingClient,
		TableId: tableId,
	}

}

func (e Employee) KickOutClients() []events.Event {
	cs := make([]client.Client, 0)

	e.club.Tables.ForEach(func(table *tables.Table) {
		c := table.Client()
		if c != nil {
			cs = append(cs, *c)
		}
	})

	sort.Slice(cs, func(i, j int) bool {
		return cs[i].Name < cs[j].Name
	})

	es := make([]events.Event, 0, len(cs))

	for _, c := range cs {
		es = append(es, events.Event{
			Time:   e.club.ClosingTime(),
			Id:     events.ClientHasGoneGen,
			Client: c,
		})
		e.EscortClientOut(c, e.club.ClosingTime())
	}

	return es
}

func (e Employee) FormTablesReport() string {
	var sb strings.Builder

	e.club.Tables.ForEach(func(table *tables.Table) {
		sb.WriteString(
			fmt.Sprintf("%d %d %s\n",
				table.Id, table.TotalSum, table.TotalTime.String()),
		)
	})

	return sb.String()

}

func (e Employee) vacateTheTable(client client.Client, time time24.Time) (int, error) {
	// хак позволяющий определить есть ли вообще клиент без использования флагов
	tableId := -1

	e.club.Tables.ForEach(func(table *tables.Table) {
		if table.HasClient(client) {
			tableId = table.Id
			table.Free()
			table.CalculateTimeAndSum(time, e.club.Tables.PricePerHour)
		}
	})

	if tableId == -1 {
		return 0, errors.New("no such client: " + client.Name)
	}
	return tableId, nil
}
