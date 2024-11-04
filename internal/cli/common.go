package cli

import (
	"log"
	"strings"

	"github.com/krzko/codemap/internal/processor"
	"github.com/urfave/cli/v2"
)

// createProcessor creates a new processor with options from CLI context
func createProcessor(c *cli.Context) (*processor.Processor, error) {
	opts := processor.DefaultOptions()
	opts.Directory = c.String("dir")
	opts.Recursive = c.Bool("recursive")
	opts.Verbose = c.Bool("verbose")

	// Parse file types
	if types := c.String("types"); types != "" {
		typeList := strings.Split(types, ",")
		for i, t := range typeList {
			t = strings.TrimSpace(t)
			if !strings.HasPrefix(t, ".") {
				typeList[i] = "." + t
			}
		}
		opts.SupportedTypes = typeList
	}

	if c.Bool("verbose") {
		log.SetFlags(log.Ltime | log.Lshortfile)
	} else {
		log.SetFlags(log.Ltime)
	}

	return processor.New(opts)
}
