package admin

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

func ListSeals(ctx context.Context, cmd *cli.Command) error {
	sealLI, err := NewCLI(cmd.String("database"), cmd.String("base-path"))
	if err != nil {
		return err
	}

	seals, err := sealLI.GetAllSeals()
	if err != nil {
		return err
	}

	fmt.Println("All Seals:")
	for _, seal := range seals {
		fmt.Printf("%+v\n", seal)
	}
	return nil
}

func ListTags(ctx context.Context, cmd *cli.Command) error {
	sealLI, err := NewCLI(cmd.String("database"), cmd.String("base-path"))
	if err != nil {
		return err
	}

	tags, err := sealLI.GetAllTags()
	if err != nil {
		return err
	}

	fmt.Println("All Tag:")
	for _, tag := range tags {
		fmt.Println(tag.Name)
	}

	return nil
}
