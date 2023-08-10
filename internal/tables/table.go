package tables

import (
	"errors"
	"fmt"
	"github.com/demig00d/computer-club/internal/client"
	"github.com/demig00d/computer-club/pkg/time24"
	"math"
)

type Table struct {
	Id        int
	client    *client.Client
	TotalSum  int
	TotalTime time24.Time
	startTime time24.Time
}

func (t *Table) CalculateTimeAndSum(endTime time24.Time, price int) {
	diff := endTime.Sub(t.startTime.Time)

	hours := int(math.Ceil(diff.Hours()))

	t.TotalSum += hours * price

	t.TotalTime = t.TotalTime.Add(diff)
}

// Вспомогательные функции, чтобы избежать nil разыменовывания поля client
func (t *Table) HasClient(client client.Client) bool {
	return t.client != nil && *t.client == client
}
func (t *Table) Free() {
	t.client = nil
}
func (t *Table) IsEmpty() bool {
	return t.client == nil
}
func (t *Table) GetClient() (client.Client, error) {
	if t.IsEmpty() {
		return client.Client{}, errors.New("table is empty")
	}

	return *t.client, nil
}

func (t *Table) SeatClient(client client.Client, time time24.Time) {
	t.client = &client
	t.startTime = time
}

func NewTable(id int) Table {
	return Table{
		Id: id,
	}
}

func CreateN(n int) []*Table {
	tables := make([]*Table, 0, n)
	for i := 1; i <= n; i++ {
		t := NewTable(i)
		tables = append(tables, &t)
	}

	return tables
}

func (t *Table) String() string {

	clientName := "<nil>"
	if t.client != nil {
		clientName = t.client.Name
	}

	return fmt.Sprintf(
		"{\n    Id: %d\n    Client: %s\n"+
			"    TotalSum: %d\n    TotalTime: %s\n"+
			"    startTime: %s\n}",
		t.Id,
		clientName,
		t.TotalSum,
		t.TotalTime.String(),
		t.startTime.String(),
	)

}
