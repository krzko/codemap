package cli

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func ListCommand() *cli.Command {
	return &cli.Command{
		Name:    "list",
		Aliases: []string{"ls"},
		Usage:   "List files that would be processed",
		Flags:   commonFlags,
		Action:  runList,
	}
}

func runList(c *cli.Context) error {
	proc, err := createProcessor(c)
	if err != nil {
		return fmt.Errorf("failed to initialize processor: %w", err)
	}

	files, err := proc.ListFiles()
	if err != nil {
		return fmt.Errorf("failed to list files: %w", err)
	}

	fmt.Printf("Found %d files in %s:\n", len(files), c.String("dir"))
	for _, file := range files {
		fmt.Println(file)
	}

	return nil
}
