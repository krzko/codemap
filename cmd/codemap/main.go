package main

import (
	"log"
	"os"
	"runtime"

	"github.com/krzko/codemap/internal/cli"
	ucli "github.com/urfave/cli/v2"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	app := &ucli.App{
		Name:    "codemap",
		Usage:   "Annotate code files with structural information for LLMs",
		Version: version,
		Metadata: map[string]interface{}{
			"commit":    commit,
			"buildDate": date,
			"goVersion": runtime.Version(),
		},
		Commands: cli.Commands(),
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
