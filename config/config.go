package config

import (
	"bufio"
	"errors"
	"github.com/demig00d/computer-club/pkg/time24"
	"strconv"
	"strings"
)

var (
	ErrTooFewTables           = errors.New("number of tables should be more than 1")
	ErrIncorrectWorkHours     = errors.New("incorrect format of work hours: should be hh:mm")
	ErrOpeningAfterClosing    = errors.New("opening time should be before closingTime")
	ErrIncorrectPricePerTable = errors.New("price per table should be more than 0")
)

type Config struct {
	openingTime  time24.Time
	closingTime  time24.Time
	pricePerHour int
	maxTables    int
}

// поля config не должны быть модифицированы вне пакета
func (c Config) OpeningTime() time24.Time {
	return c.openingTime
}
func (c Config) ClosingTime() time24.Time {
	return c.closingTime
}
func (c Config) PricePerHour() int {
	return c.pricePerHour

}
func (c Config) MaxTables() int {
	return c.maxTables
}

func NewConfig(fileScanner *bufio.Scanner) (Config, error) {
	cfg := Config{}

	i := 0
	for i < 3 && fileScanner.Scan() {
		err := cfg.getLine(i, fileScanner.Text())
		if err != nil {
			return Config{}, err
		}
		i++
	}

	if i < 3 {
		return Config{}, errors.New("not enough lines of configuration")
	}

	return cfg, nil
}

func (c *Config) getLine(i int, line string) error {
	switch i {
	case 0:
		n, err := strconv.Atoi(line)
		if n < 1 {
			return ErrTooFewTables
		}

		if err != nil {
			return err
		}

		c.maxTables = n

	case 1:
		fields := strings.Fields(line)
		if len(fields) != 2 {
			return ErrIncorrectWorkHours
		}

		startTime, err := time24.Parse(fields[0])
		endTime, err := time24.Parse(fields[1])
		if err != nil {
			return err
		}

		if !endTime.After(startTime.Time) {
			return ErrOpeningAfterClosing
		}

		c.openingTime = startTime
		c.closingTime = endTime

	case 2:
		n, err := strconv.Atoi(line)
		if n < 1 {
			return ErrIncorrectPricePerTable
		}

		if err != nil {
			return err
		}

		c.pricePerHour = n
	}

	return nil
}
