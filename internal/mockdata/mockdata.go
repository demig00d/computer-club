package mockdata

import (
	"bufio"
	"github.com/demig00d/computer-club/config"
	"github.com/demig00d/computer-club/internal/client"

	"github.com/demig00d/computer-club/pkg/time24"
	"strings"
)

var (
	Time00, _ = time24.Parse("00:00")
	Time02, _ = time24.Parse("02:00")
	Time09, _ = time24.Parse("09:00")
	Time11, _ = time24.Parse("11:00")
	Time19, _ = time24.Parse("19:00")
	Client1   = client.Client{Name: "client1"}
)

func MockConfig() config.Config {
	cfg, _ := config.NewConfig(
		bufio.NewScanner(strings.NewReader("5\n10:00 20:00\n10")),
	)
	return cfg
}
