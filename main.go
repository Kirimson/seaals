package main

import (
	"context"
	"errors"
	"log"
	"os"
	"seaals/cmd"
	"seaals/cmd/admin"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "seaals",
		Usage: "Of Approval",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "database",
				Aliases: []string{"db"},
				Value:   "seaals.db",
			},
		},
		Commands: []*cli.Command{
			{
				Name:   "serve",
				Usage:  "Run the SEAaLS API Server",
				Action: cmd.Serve,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "base-path",
						Aliases: []string{"b"},
						Value:   "./",
						Validator: func(v string) error {
							// Ensure the path provided exists
							if _, err := os.Stat(v); err != nil {
								return err
							}
							return nil
						},
					},
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
			{
				Name:  "admin",
				Usage: "Perform administrative seal activities",
				Commands: []*cli.Command{
					{
						Name:   "list-seals",
						Usage:  "List all available seals",
						Action: admin.ListSeals,
					},
					{
						Name:   "list-tags",
						Usage:  "List all available seals",
						Action: admin.ListTags,
					},
					{
						Name:  "add-seal",
						Usage: "Add a new seal to the available catalogue",
						Arguments: []cli.Argument{
							&cli.StringArg{
								Name: "path",
							},
						},
						Flags: []cli.Flag{
							&cli.StringSliceFlag{
								Name:  "tag",
								Usage: "Extra tags to associate with the Seal",
							},
						},
						Action: admin.AddSeal,
					},
					{
						Name:  "add-tag",
						Usage: "Add a new tag to the available catalogue",
						Arguments: []cli.Argument{
							&cli.StringArg{
								Name: "name",
							},
						},
						Action: admin.AddTag,
					},
				},
			},
			{
				Name:  "db",
				Usage: "Perform Database actions",
				Commands: []*cli.Command{
					{
						Name:   "migrate",
						Usage:  "Migrate the SEAaLS database",
						Action: cmd.Migrate,
					},
				},
			},
		},
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
