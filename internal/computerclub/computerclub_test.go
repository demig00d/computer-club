package computerclub

import (
	"bufio"
	"github.com/demig00d/computer-club/config"
	"github.com/demig00d/computer-club/internal/client"
	"github.com/demig00d/computer-club/internal/tables"
	"github.com/demig00d/computer-club/pkg/queue"
	"github.com/demig00d/computer-club/pkg/time24"
	"strings"
	"testing"
)

var time00, _ = time24.Parse("00:00")
var mockCfg, _ = config.NewConfig(
	bufio.NewScanner(strings.NewReader("5\n10:00 20:00\n10")),
)

var mockComputerClub = NewComputerClub(mockCfg)

func ClientsMoreThanTablesClub() ComputerClub {
	club := NewComputerClub(mockCfg)

	cs := []client.Client{
		{"client1"},
		{"client2"},
		{"client3"},
		{"client4"},
		{"client5"},
		{"client6"},
		{"client7"},
	}

	q := queue.NewQueue[client.Client]()

	for _, c := range cs {
		q.Push(c)
	}

	club.ClientQueue = q

	return club
}

func TestComputerClub_AreClientsInQueueMoreThanTables(t *testing.T) {
	club := ClientsMoreThanTablesClub()

	if !club.AreClientsInQueueMoreThanTables() {
		t.Errorf("returned false, but available tables=%d < clients in queue=%d",
			len(club.Tables), club.ClientQueue.Length())
	}

	club.ClientQueue.Poll()
	club.ClientQueue.Poll()

	if club.AreClientsInQueueMoreThanTables() {
		t.Errorf("returned true, but available tables=%d > clients in queue=%d",
			len(club.Tables), club.ClientQueue.Length())
	}

}

func TestComputerClub_IsBeforeOpening(t *testing.T) {
	club := mockComputerClub

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
	club := mockComputerClub

	client1 := client.Client{"client1"}

	if club.IsClientIn(client1) {
		t.Errorf("returned true, but client with name '%s' is not in club:\n    club.Tables: %v",
			client1.Name, club.Tables)
	}

	club.Tables[0].SeatClient(client1, time00)

	if !club.IsClientIn(client1) {
		t.Errorf("returned false, but client with name '%s' is in club:\n    club.Tables: %v",
			client1.Name, club.Tables)
	}
}

func TestComputerClub_AreThereAnyTablesAvailable(t *testing.T) {
	club := mockComputerClub

	if !club.AreThereAnyTablesAvailable() {
		t.Errorf("returned false, but there are available tables: %v",
			club.Tables)
	}

	club.Tables = []*tables.Table{}

	if club.AreThereAnyTablesAvailable() {
		t.Errorf("returned true, but there are no tables available: %v",
			club.Tables)
	}

}
