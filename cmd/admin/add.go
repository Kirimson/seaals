package admin

import (
	"context"
	"fmt"
	"os"

	"github.com/urfave/cli/v3"
)

func AddSeal(ctx context.Context, cmd *cli.Command) error {
	sealLI, err := NewCLI(cmd.String("database"), cmd.String("base-path"))
	if err != nil {
		return err
	}

	// Try and read the local file
	sealData, err := os.ReadFile(cmd.StringArg(("path")))
	if err != nil {
		return err
	}

	// Add this image as a new Seal
	newSeal, err := sealLI.AddSeal(sealData, cmd.StringSlice("tag"))
	if err != nil {
		return err
	}

	fmt.Println("New Seal!")
	fmt.Println(newSeal.Path)
	return nil
}

func AddTag(ctx context.Context, cmd *cli.Command) error {
	sealLI, err := NewCLI(cmd.String("database"), cmd.String("base-path"))
	if err != nil {
		return err
	}

	newTag, err := sealLI.AddTag(cmd.StringArg("name"))
	if err != nil {
		return err
	}

	fmt.Println("New Tag!")
	fmt.Println(newTag.Name)
	return nil
}
