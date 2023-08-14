package computerclub

import (
	"github.com/demig00d/computer-club/internal/client"
	"github.com/demig00d/computer-club/internal/mockdata"
	"github.com/demig00d/computer-club/internal/tables"
	"github.com/demig00d/computer-club/pkg/queue"
	"github.com/demig00d/computer-club/pkg/time24"
	"testing"
)

func MockComputerClub() ComputerClub {
	return NewComputerClub(mockdata.MockConfig())
}

func ClientsMoreThanTablesClub() ComputerClub {
	club := MockComputerClub()

	cs := []client.Client{
		{Name: "client1"},
		{Name: "client2"},
		{Name: "client3"},
		{Name: "client4"},
		{Name: "client5"},
		{Name: "client6"},
		{Name: "client7"},
	}

	q := queue.NewQueue[client.Client]()

	for _, c := range cs {
		_ = q.Push(c)
	}

	club.ClientQueue = q

	return club
}

func TestComputerClub_AreClientsInQueueMoreThanTables(t *testing.T) {
	club := ClientsMoreThanTablesClub()

	if !club.AreClientsInQueueMoreThanTables() {
		t.Errorf("returned false, but true expected")
	}

	_, _ = club.ClientQueue.Poll()
	_, _ = club.ClientQueue.Poll()

	if club.AreClientsInQueueMoreThanTables() {
		t.Errorf("returned true, but false expected")
	}

}

func TestComputerClub_IsBeforeOpening(t *testing.T) {
	club := MockComputerClub()

	timeBeforeOpening, _ := time24.Parse("02:00")
	timeAfterOpening, _ := time24.Parse("22:00")

	if !club.IsBeforeOpening(timeBeforeOpening) {
		t.Errorf("returned false, but time %s is before opening time %s",
			timeBeforeOpening, club.OpeningTime())
	}

	if club.IsBeforeOpening(timeAfterOpening) {
		t.Errorf("returned true, but time %s is not before opening time %s",
			timeAfterOpening, club.OpeningTime())
	}
}

func TestComputerClub_IsClientIn(t *testing.T) {
	club := MockComputerClub()

	client1 := mockdata.Client1

	if club.IsClientIn(client1) {
		t.Errorf("returned true, but client with name '%s' is not in club:\n    club.Tables: %v",
			client1.Name, club.Tables)
	}

	club.Tables.ForEach(func(table *tables.Table) {
		if table.Id == 1 {
			table.SeatClient(client1, mockdata.Time00)
		}
	})

	if !club.IsClientIn(client1) {
		t.Errorf("returned false, but client with name '%s' is in club:\n    club.Tables: %v",
			client1.Name, club.Tables)
	}
}

func TestComputerClub_AreThereAnyTablesAvailable(t *testing.T) {
	club := MockComputerClub()
	client1 := mockdata.Client1

	if !club.AreThereAnyTablesAvailable() {
		t.Errorf("returned false, but there are available tables: %v",
			club.Tables)
	}

	club.Tables.ForEach(func(table *tables.Table) {
		table.SeatClient(client1, mockdata.Time00)
	})

	if club.AreThereAnyTablesAvailable() {
		t.Errorf("returned true, but there are no tables available: %v",
			club.Tables)
	}

}
