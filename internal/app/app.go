package app

import (
	"bufio"
	"fmt"
	"github.com/demig00d/computer-club/config"
	"github.com/demig00d/computer-club/internal/events"
	"github.com/demig00d/computer-club/internal/staff"
	"log"
)

type App struct {
	manager     staff.Manager
	fileScanner *bufio.Scanner
}

func NewApp(cfg config.Config, filescanner *bufio.Scanner) App {
	return App{
		manager:     staff.NewManager(cfg),
		fileScanner: filescanner,
	}
}

func Run(app App) {

	app.manager.OpenClub()

	for app.fileScanner.Scan() {
		line := app.fileScanner.Text()
		event, err := events.Parse(line)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(event)

		newEvent := app.manager.HandleEvent(event)

		if newEvent != nil {
			fmt.Println(newEvent)
		}
	}

	app.manager.CloseClub()
}
