package admin

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/urfave/cli/v3"
)

func AddSeal(ctx context.Context, cmd *cli.Command) error {
	sealLI, err := NewCLI(cmd.String("database"), cmd.String("base-path"))
	if err != nil {
		return err
	}

	var sealData []byte
	if cmd.String("path") != "" {
		// Try and read the local file
		sealData, err = sealFromFile(cmd.String("path"))
		if err != nil {
			return err
		}
	} else if cmd.String("url") != "" {
		// Try and download from the URL
		sealData, err = sealFromURL(cmd.String("url"))
		if err != nil {
			return err
		}
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

func sealFromFile(path string) ([]byte, error) {
	fmt.Println("Reading seal...")
	return os.ReadFile(path)
}

func sealFromURL(url string) ([]byte, error) {
	fmt.Printf("Downloading seal from '%s'...\n", url)
	client := http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Accept", "image/*")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	// Ensure downloaded file is an image
	fmt.Println(resp.Header.Get("Content-Type"))
	if strings.HasPrefix(resp.Header.Get("Content-Type"), "image/") {
		sealData, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return sealData, nil
	}
	return nil, errors.New("URL does not point to an image")
}

func AddTag(ctx context.Context, cmd *cli.Command) error {
	sealLI, err := NewCLI(cmd.String("database"), cmd.String("base-path"))
	if err != nil {
		return err
	}

	tagName := cmd.StringArg("name")
	if tagName == "" {
		return errors.New("no tag name provided")
	}

	newTag, err := sealLI.AddTag(tagName)
	if err != nil {
		return err
	}

	fmt.Println("New Tag!")
	fmt.Println(newTag.Name)
	return nil
}
