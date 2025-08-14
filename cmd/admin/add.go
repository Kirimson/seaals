package admin

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

func Add(ctx context.Context, cmd *cli.Command) error {
	sealLI, err := NewCLI(cmd.String("database"), cmd.String("base-path"))
	if err != nil {
		return err
	}

	newSeal, err := sealLI.AddSeal(cmd.StringArg("path"), cmd.StringSlice("tag"))
	if err != nil {
		return err
	}

	fmt.Println("New Seal!")
	fmt.Println(newSeal.Path)

	return nil
}
