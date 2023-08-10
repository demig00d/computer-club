package config

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/demig00d/computer-club/pkg/time24"
	"strings"
	"testing"
	"time"
)

var validConfigs = []Config{
	{
		openingTime:  genHours(9),
		closingTime:  genHours(19),
		pricePerHour: 10,
		maxTables:    3,
	},
	{
		openingTime:  genHours(10),
		closingTime:  genHours(20),
		pricePerHour: 5,
		maxTables:    1,
	},
	{
		openingTime:  genHours(8),
		closingTime:  genHours(15),
		pricePerHour: 99,
		maxTables:    4,
	},
}

var InvalidConfigsTable = []struct {
	configString string
	expectedErr  error
}{
	{
		configString: "0\n14:00 12:00\n12",
		expectedErr:  ErrTooFewTables,
	},
	{
		configString: "5\n14:00\n12",
		expectedErr:  ErrIncorrectWorkHours,
	},
	{
		configString: "5\n14:00 9:00\n12",
		expectedErr:  ErrOpeningAfterClosing,
	},
	{
		configString: "3\n14:00 19:00\n0",
		expectedErr:  ErrIncorrectPricePerTable,
	},
}

func TestConfig(t *testing.T) {

	for _, expected := range validConfigs {
		scanner := getScanner(expected)
		got, err := NewConfig(scanner)

		if err != nil {
			t.Errorf("unexpected error %s", err.Error())
		}

		if got != (expected) {
			t.Errorf("got:\n%v\n\nexpected:\n%v", got, expected)
		}
	}

	for _, cfg := range InvalidConfigsTable {
		r := strings.NewReader(cfg.configString)
		scanner := bufio.NewScanner(r)
		_, err := NewConfig(scanner)

		if !errors.Is(err, cfg.expectedErr) {
			t.Errorf("got: %s, expected: %s", err, cfg.expectedErr)
		}
	}
}

// utils
var time00, _ = time24.Parse("00:00")

func genHours(h time.Duration) time24.Time {
	return time00.Add(h * time.Hour)
}

func getScanner(cfg Config) *bufio.Scanner {
	configString := fmt.Sprintf("%d\n%s %s\n%d",
		cfg.maxTables,
		cfg.openingTime.String(),
		cfg.closingTime.String(),
		cfg.pricePerHour,
	)
	r := strings.NewReader(configString)
	return bufio.NewScanner(r)
}

func (cfg Config) String() string {
	return fmt.Sprintf("%d\n%s %s\n%d",
		cfg.maxTables,
		cfg.openingTime.String(),
		cfg.closingTime.String(),
		cfg.pricePerHour,
	)
}
