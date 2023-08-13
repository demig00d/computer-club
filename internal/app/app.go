package app

import (
	"bufio"
	"fmt"
	"github.com/demig00d/computer-club/config"
	"github.com/demig00d/computer-club/internal/computerclub"
	"github.com/demig00d/computer-club/internal/events"
	"github.com/demig00d/computer-club/internal/staff"
	"log"
)

type App struct {
	config      config.Config
	fileScanner *bufio.Scanner
}

func NewApp(cfg config.Config, filescanner *bufio.Scanner) App {
	return App{
		config:      cfg,
		fileScanner: filescanner,
	}
}

func Run(app App) {

	club := computerclub.NewComputerClub(app.config)
	employee := staff.NewEmployee(club)
	manager := staff.NewManager(club.Workhours, employee)

	manager.OpenClub()

	for app.fileScanner.Scan() {
		line := app.fileScanner.Text()
		event, err := events.Parse(line)
		if err != nil {
			log.Fatal(err)
		}

		pEvent := &event
		for pEvent != nil {
			fmt.Println(pEvent)
			pEvent = manager.HandleEvent(*pEvent)
		}

	}

	manager.CloseClub()
}
