package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/krzko/codemap/internal/processor"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func init() {
	log.SetFlags(log.Ltime)
}

func displayVersion() {
	fmt.Printf("codemap version %s\n", version)
	fmt.Printf("  Build date: %s\n", date)
	fmt.Printf("  Git commit: %s\n", commit)
	fmt.Printf("  Go version: %s\n", runtime.Version())
	fmt.Printf("  OS/Arch: %s/%s\n", runtime.GOOS, runtime.GOARCH)
}

func main() {
	showVersion := flag.Bool("version", false, "Display version information")

	dir := flag.String("dir", ".", "Directory to process (defaults to current directory)")
	clean := flag.Bool("clean", false, "Remove existing annotations")
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	fileTypes := flag.String("types", "go,py,js,jsx,ts,tsx", "Comma-separated list of file types to process")

	versionShort := flag.Bool("v", false, "Display version information")

	flag.Parse()

	// Check for version flag
	if *showVersion || *versionShort {
		displayVersion()
		os.Exit(0)
	}

	if *verbose {
		log.SetFlags(log.Ltime | log.Lshortfile)
	}

	absPath, err := filepath.Abs(*dir)
	if err != nil {
		log.Printf("Warning: Could not resolve absolute path for %s: %v", *dir, err)
		absPath = *dir
	}

	log.Printf("Processing directory: %s", absPath)
	log.Printf("Looking for file types: %s", *fileTypes)

	opts := processor.DefaultOptions()
	opts.Directory = *dir
	opts.Clean = *clean

	proc, err := processor.New(opts)
	if err != nil {
		log.Fatalf("Failed to initialise processor: %v", err)
	}

	if err := proc.Process(); err != nil {
		log.Fatalf("Failed to process directory: %v", err)
		os.Exit(1)
	}

	log.Println("Processing completed successfully")
}
