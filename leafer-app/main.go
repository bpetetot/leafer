package main

import (
	"log"
	"os"

	"github.com/bpetetot/leafer/db"
	"github.com/bpetetot/leafer/server"
	"github.com/bpetetot/leafer/services"
	"github.com/bpetetot/leafer/services/upnp"
	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
)

func main() {
	// Load dotenv file
	if godotenv.Load() != nil {
		log.Fatal("Error loading .env file.")
	}

	// Initiate CLI
	app := cli.NewApp()

	app.Name = "leafer"
	app.HelpName = "leafer"
	app.Usage = "your media center for libraries."
	app.EnableBashCompletion = true

	app.Commands = []*cli.Command{
		{
			Name:  "serve",
			Usage: "start Leafer server",
			Action: func(c *cli.Context) error {
				server.Start()
				return nil
			},
		}, {
			Name:  "expose",
			Usage: "expose server through UPNP router",
			Action: func(c *cli.Context) error {
				upnp.Expose()
				return nil
			},
		}, {
			Name:  "unexpose",
			Usage: "Unexpose server through UPNP router",
			Action: func(c *cli.Context) error {
				upnp.Unexpose()
				return nil
			},
		},
		{
			Name:  "analyze",
			Usage: "start a library analysis",
			Action: func(c *cli.Context) error {
				DB := db.Setup()
				scanner := services.NewScannerService(DB)
				scraper := services.NewScraperService(DB)

				err := scanner.ScanLibrary(1)
				if err != nil {
					return err
				}

				err = scraper.ScrapLibrary(1)
				if err != nil {
					return err
				}
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
