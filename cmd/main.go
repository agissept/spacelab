package main

import (
	"log"
	"os"

	"github.com/agissept/spacelab/internal"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "spacelab",
		Usage: "lab for make starship.rs more fun",
		Commands: []*cli.Command{
			{
				Name:    "now-playing",
				Aliases: []string{"np"},
				Usage:   "Print now playing on spotify",
				Action: func(ctx *cli.Context) error {
					internal.NowPlaying()
					return nil
				},
			},
			{
				Name:    "play-poker",
				Aliases: []string{"pp"},
				Usage:   "Played hands",
				Action: func(ctx *cli.Context) error {
					internal.PlayPoker()
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}