package cli

import (
	"fmt"
	"log"

	"github.com/urfave/cli/v2"
)

func CleanCommand() *cli.Command {
	return &cli.Command{
		Name:  "clean",
		Usage: "Remove annotations from files",
		Flags: append(commonFlags,
			&cli.BoolFlag{
				Name:    "dry-run",
				Aliases: []string{"n"},
				Usage:   "Show what would be done without making changes",
			},
		),
		Action: runClean,
	}
}

func runClean(c *cli.Context) error {
	proc, err := createProcessor(c)
	if err != nil {
		return fmt.Errorf("failed to initialize processor: %w", err)
	}

	files, err := proc.ListFiles()
	if err != nil {
		return fmt.Errorf("failed to list files: %w", err)
	}

	if c.Bool("dry-run") {
		log.Printf("Would clean %d files in %s", len(files), c.String("dir"))
		for _, file := range files {
			log.Printf("Would clean: %s", file)
		}
		return nil
	}

	if err := proc.Clean(); err != nil {
		return fmt.Errorf("failed to clean files: %w", err)
	}

	return nil
}
