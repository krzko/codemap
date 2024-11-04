package cli

import (
	"fmt"
	"log"

	"github.com/urfave/cli/v2"
)

func ApplyCommand() *cli.Command {
	return &cli.Command{
		Name:  "apply",
		Usage: "Add annotations to files",
		Flags: append(commonFlags,
			&cli.BoolFlag{
				Name:    "dry-run",
				Aliases: []string{"n"},
				Usage:   "Show what would be done without making changes",
			},
			&cli.BoolFlag{
				Name:    "recursive",
				Aliases: []string{"r"},
				Usage:   "Process directories recursively",
				Value:   true,
			},
		),
		Action: runApply,
	}
}

func runApply(c *cli.Context) error {
	proc, err := createProcessor(c)
	if err != nil {
		return fmt.Errorf("failed to initialize processor: %w", err)
	}

	files, err := proc.ListFiles()
	if err != nil {
		return fmt.Errorf("failed to list files: %w", err)
	}

	if c.Bool("dry-run") {
		log.Printf("Would process %d files in %s", len(files), c.String("dir"))
		for _, file := range files {
			log.Printf("Would annotate: %s", file)
		}
		return nil
	}

	if err := proc.Process(); err != nil {
		return fmt.Errorf("failed to process files: %w", err)
	}

	return nil
}