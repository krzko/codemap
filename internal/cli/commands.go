package cli

import (
	"github.com/urfave/cli/v2"
)

func Commands() []*cli.Command {
	return []*cli.Command{
		ApplyCommand(),
		CleanCommand(),
		ListCommand(),
		StatsCommand(),
	}
}

var commonFlags = []cli.Flag{
	&cli.StringFlag{
		Name:    "dir",
		Aliases: []string{"d"},
		Usage:   "Directory to process",
		Value:   ".",
	},
	&cli.StringFlag{
		Name:    "types",
		Aliases: []string{"t"},
		Usage:   "Comma-separated list of file types to process",
		Value:   "go,py,js,jsx,ts,tsx",
	},
	&cli.BoolFlag{
		Name:    "verbose",
		Aliases: []string{"V"},
		Usage:   "Enable verbose logging",
	},
}
