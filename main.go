package main

import (
	"context"
	"errors"
	"log"
	"os"
	"seaals/cmd"
	"seaals/cmd/admin"

	"github.com/joho/godotenv"
	"github.com/urfave/cli/v3"
)

func main() {
	// Load env vars
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cmd := &cli.Command{
		Name:  "seaals",
		Usage: "Of Approval",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "database",
				Aliases: []string{"db"},
				Value:   "seaals.db",
				Sources: cli.EnvVars("DB_PATH"),
			},
			&cli.StringFlag{
				Name:    "base-path",
				Aliases: []string{"b"},
				Value:   "./",
				Sources: cli.EnvVars("BASE_PATH"),
				Validator: func(v string) error {
					// Ensure the path provided exists
					if _, err := os.Stat(v); err != nil {
						return err
					}
					return nil
				},
			},
		},
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
						Value:   8080,
						Usage:   "Port to listen on. Valid ports: 1 - 65535",
						Sources: cli.EnvVars("PORT"),
					},
				},
			},
			{
				Name:  "admin",
				Usage: "Perform administrative seal activities",
				Commands: []*cli.Command{
					{
						Name:  "list-seals",
						Usage: "List all available seals",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name: "tag",
							},
						},
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
						MutuallyExclusiveFlags: []cli.MutuallyExclusiveFlags{
							{
								Required: true,
								Flags: [][]cli.Flag{
									{
										&cli.StringFlag{
											Name: "path",
										},
									},
									{
										&cli.StringFlag{
											Name: "url",
										},
									},
								},
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
