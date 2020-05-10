package main

import (
	"log"
	"os"

	"github.com/bpetetot/leafer/db"
	"github.com/bpetetot/leafer/scanners"
	"github.com/bpetetot/leafer/server"
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
			Name:  "server",
			Usage: "start Leafer server",
			Action: func(c *cli.Context) error {
				server.Start()
				return nil
			},
		},
		{
			Name:  "analyze",
			Usage: "start a library analysis",
			Action: func(c *cli.Context) error {
				conn := db.Setup()
				library := db.Library{Name: "My mangas", Path: c.Args().First()}
				conn.Create(&library)
				// var library db.Library
				// conn.First(&library)

				scanners.ScanFiles(&library, conn)
				scanners.ScanMedias(&library, conn)
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
