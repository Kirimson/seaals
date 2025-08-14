package admin

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

func List(ctx context.Context, cmd *cli.Command) error {
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
		fmt.Println(seal.Path)
	}

	return nil
}
