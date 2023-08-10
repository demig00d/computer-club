package computerclub

import (
	"github.com/demig00d/computer-club/config"
	"github.com/demig00d/computer-club/internal/client"
	"github.com/demig00d/computer-club/internal/tables"
	"github.com/demig00d/computer-club/pkg/queue"
	"github.com/demig00d/computer-club/pkg/time24"
)

// ComputerClub хранит состояние приложения
// и имеет методы для проверки
type ComputerClub struct {
	config.Config
	ClientQueue queue.Queue[client.Client]
	Tables      []*tables.Table
}

func NewComputerClub(cfg config.Config) ComputerClub {
	return ComputerClub{
		Config:      cfg,
		ClientQueue: queue.NewQueue[client.Client](),
		Tables:      tables.CreateN(cfg.MaxTables()),
	}
}

func (c ComputerClub) IsBeforeOpening(time time24.Time) bool {
	return c.OpeningTime().After(time.Time)
}

func (c ComputerClub) AreThereAnyTablesAvailable() bool {
	for _, t := range c.Tables {
		if t.IsEmpty() {
			return true
		}
	}
	return false
}

func (c ComputerClub) AreClientsInQueueMoreThanTables() bool {
	availableTablesCount := 0
	for _, table := range c.Tables {
		if table.IsEmpty() {
			availableTablesCount++
		}
	}
	return c.ClientQueue.Length() > availableTablesCount
}

func (c ComputerClub) IsClientIn(client client.Client) bool {

	isInTable := func() bool {
		for _, t := range c.Tables {
			if t.HasClient(client) {
				return true
			}
		}
		return false
	}

	return c.ClientQueue.IsIn(client) || isInTable()

}
