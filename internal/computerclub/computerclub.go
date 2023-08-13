package computerclub

import (
	"github.com/demig00d/computer-club/config"
	"github.com/demig00d/computer-club/internal/client"
	"github.com/demig00d/computer-club/internal/tables"
	"github.com/demig00d/computer-club/internal/workhours"
	"github.com/demig00d/computer-club/pkg/queue"
	"github.com/demig00d/computer-club/pkg/time24"
)

// ComputerClub хранит состояние приложения
// и имеет методы для проверки
type ComputerClub struct {
	workhours.Workhours
	ClientQueue queue.Queue[client.Client]
	Tables      tables.Tables
}

func NewComputerClub(cfg config.Config) ComputerClub {

	workhours := workhours.NewWorkhours(cfg)

	return ComputerClub{
		Workhours:   workhours,
		ClientQueue: queue.NewQueue[client.Client](),
		Tables:      tables.NewTables(cfg.PricePerHour(), cfg.MaxTables()),
	}
}

func (c ComputerClub) AreThereAnyTablesAvailable() bool {
	areAnyAvailable := false

	c.Tables.ForEach(func(t *tables.Table) {
		if t.IsEmpty() {
			areAnyAvailable = true
		}
	})

	return areAnyAvailable
}

func (c ComputerClub) AreClientsInQueueMoreThanTables() bool {
	availableTablesCount := 0
	c.Tables.ForEach(func(table *tables.Table) {
		if table.IsEmpty() {
			availableTablesCount++
		}
	})
	return c.ClientQueue.Length() > availableTablesCount
}

func (c ComputerClub) IsClientIn(client client.Client) bool {

	isInTable := func() bool {
		inTable := false
		c.Tables.ForEach(func(t *tables.Table) {
			if t.HasClient(client) {
				inTable = true
			}
		})
		return inTable
	}

	return c.ClientQueue.IsIn(client) || isInTable()

}

func (c ComputerClub) IsBeforeOpening(time time24.Time) bool {
	return c.OpeningTime().After(time.Time)
}
