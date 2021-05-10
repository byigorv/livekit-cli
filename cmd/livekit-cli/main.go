package main

import (
	"fmt"
	"os"

	livekit_cli "github.com/livekit/livekit-cli"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "livekit-cli",
		Usage: "CLI client to LiveKit",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name: "verbose",
			},
		},
		Version: livekit_cli.Version,
	}

	app.Commands = append(app.Commands, TokenCommands...)
	app.Commands = append(app.Commands, RoomCommands...)
	app.Commands = append(app.Commands, JoinCommands...)

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}