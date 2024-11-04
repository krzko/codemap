package cli

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func StatsCommand() *cli.Command {
	return &cli.Command{
		Name:   "stats",
		Usage:  "Show statistics about annotations",
		Flags:  commonFlags,
		Action: runStats,
	}
}

func runStats(c *cli.Context) error {
	proc, err := createProcessor(c)
	if err != nil {
		return fmt.Errorf("failed to initialize processor: %w", err)
	}

	stats, err := proc.GetStats()
	if err != nil {
		return fmt.Errorf("failed to get statistics: %w", err)
	}

	// Print statistics
	fmt.Printf("Statistics for %s:\n", c.String("dir"))
	fmt.Printf("Total files processed: %d\n", stats.TotalFiles)
	fmt.Printf("Files with annotations: %d\n", stats.AnnotatedFiles)
	fmt.Printf("Files without annotations: %d\n", stats.UnannotatedFiles)

	fmt.Println("\nBreakdown by language:")
	for lang, count := range stats.FilesByLanguage {
		fmt.Printf("  %s: %d files\n", lang, count)
	}

	return nil
}
