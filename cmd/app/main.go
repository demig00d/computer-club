package main

import (
	"bufio"
	"log"
	"os"

	"github.com/demig00d/computer-club/config"
	"github.com/demig00d/computer-club/internal/app"
)

func main() {
	filename := os.Args[1]

	readFile, err := os.Open(filename)
	defer func() {
		if err := readFile.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	cfg, err := config.NewConfig(fileScanner)
	if err != nil {
		log.Fatal(err)
	}

	application := app.NewApp(cfg, fileScanner)

	app.Run(application)
}
