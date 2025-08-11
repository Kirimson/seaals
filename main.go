package main

import (
	"context"
	"errors"
	"log"
	"os"
	"seaals/cmd"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "seaals",
		Usage: "Of Approval",
		Commands: []*cli.Command{
			{
				Name:   "serve",
				Usage:  "Run the SEAaLS API Server",
				Action: cmd.Serve,
				Flags: []cli.Flag{
					&cli.Int16Flag{
						Name:    "port",
						Aliases: []string{"p"},
						Validator: func(v int16) error {
							if v == 0 {
								return errors.New("port cannot be 0")
							}
							return nil
						},
						Value: 8080,
						Usage: "Port to listen on. Valid ports: 1 - 65535",
					},
				},
			},
		},
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
