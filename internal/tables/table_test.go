package tables

import (
	"github.com/demig00d/computer-club/internal/mockdata"
	"testing"
)

var table = Table{
	Id:        1,
	client:    nil,
	TotalSum:  0,
	TotalTime: mockdata.Time00,
	startTime: mockdata.Time00,
}

func TestTable_CalculateTimeAndSum(t *testing.T) {
	table.CalculateTimeAndSum(mockdata.Time19, 10)

	expectedTotalSum := 190
	expectedTotalTime := mockdata.Time19

	if expectedTotalSum != table.TotalSum {
		t.Errorf("got: %d, expected: %d", table.TotalSum, expectedTotalSum)
	}
	if expectedTotalTime != table.TotalTime {
		t.Errorf("got: %v, expected: %v", table.TotalTime, expectedTotalTime)
	}

}

func TestTable_HasClient(t *testing.T) {

	if table.HasClient(mockdata.Client1) {
		t.Errorf("returned true, but table is empty: %v", &table)
	}

	table.client = &mockdata.Client1

	if !table.HasClient(mockdata.Client1) {
		t.Errorf("returned false, but table client belongs to table: %v", &table)
	}
}

func TestTable_SeatClient(t *testing.T) {
	table.client = nil

	table.SeatClient(mockdata.Client1, mockdata.Time00)

	if table.client == nil || *table.client != mockdata.Client1 {
		t.Errorf("got %s, but table.client should be equal %s",
			table.client, &mockdata.Client1)
	}

}

func TestTable_Free(t *testing.T) {
	table.client = &mockdata.Client1
	table.Free()

	if table.client != nil {
		t.Errorf("got %s, but table.client should be nil",
			table.client)
	}
}
