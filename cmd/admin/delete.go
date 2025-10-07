package admin

import (
	"context"
	"errors"
	"fmt"

	"github.com/urfave/cli/v3"
)

func DeleteSeal(ctx context.Context, cmd *cli.Command) error {
	sealLI, err := NewCLI(cmd.String("database"), cmd.String("base-path"))
	if err != nil {
		return err
	}

	sealID := cmd.Int64Arg("id")
	if sealID == 0 {
		return errors.New("no seal ID provided")
	}
	if err := sealLI.DeleteSeal(sealID); err != nil {
		return err
	}
	fmt.Printf("Deleted seal with ID %d\n", sealID)
	return nil
}
