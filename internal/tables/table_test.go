package tables

import (
	"errors"
	"github.com/demig00d/computer-club/internal/client"
	"github.com/demig00d/computer-club/pkg/time24"
	"testing"
)

var time00, _ = time24.Parse("00:00")
var time19, _ = time24.Parse("19:00")
var client1 = client.Client{"client1"}

var table = Table{
	Id:        1,
	client:    nil,
	TotalSum:  0,
	TotalTime: time00,
	startTime: time00,
}

func TestTable_CalculateTimeAndSum(t *testing.T) {
	table.CalculateTimeAndSum(time19, 10)

	expectedTotalSum := 190
	expectedTotalTime := time19

	if expectedTotalSum != table.TotalSum {
		t.Errorf("got: %d, expected: %d", table.TotalSum, expectedTotalSum)
	}
	if expectedTotalTime != table.TotalTime {
		t.Errorf("got: %v, expected: %v", table.TotalTime, expectedTotalTime)
	}

}

func TestTable_HasClient(t *testing.T) {

	if table.HasClient(client1) {
		t.Errorf("returned true, but table is empty: %v", &table)
	}

	table.client = &client1

	if !table.HasClient(client1) {
		t.Errorf("returned false, but table client belongs to table: %v", &table)
	}
}

func TestTable_SeatClient(t *testing.T) {
	table.client = nil

	table.SeatClient(client1, time00)

	if table.client == nil || *table.client != client1 {
		t.Errorf("got %s, but table.client should be equal %s",
			table.client, &client1)
	}

}

func TestTable_Free(t *testing.T) {
	table.client = &client1
	table.Free()

	if table.client != nil {
		t.Errorf("got %s, but table.client should be nil",
			table.client)
	}
}

func TestTable_GetClient(t *testing.T) {
	table.client = &client1

	c, err := table.GetClient()

	if err != nil {
		t.Errorf("unexpected error %s", err.Error())
	}

	if c != client1 {
		t.Errorf("clients mismamtch, got: %v != %v", c, client1)
	}

	table.client = nil

	_, err = table.GetClient()
	if !errors.Is(err, ErrEmptyTable) {
		t.Error("got nil error, but ErrEmptyTable expected")
	}

}
