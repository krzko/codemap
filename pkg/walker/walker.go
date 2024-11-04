package walker

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gobwas/glob"
)

type Walker struct {
	root         string
	excludeDirs  []string
	excludeFiles []glob.Glob
}

func New(root string, opts ...Option) (*Walker, error) {
	// If root is empty or ".", use the current working directory
	if root == "" || root == "." {
		wd, err := os.Getwd()
		if err != nil {
			return nil, fmt.Errorf("failed to get working directory: %w", err)
		}
		root = wd
	}

	// Ensure the root path is absolute
	absRoot, err := filepath.Abs(root)
	if err != nil {
		return nil, fmt.Errorf("failed to get absolute path for root directory: %w", err)
	}

	w := &Walker{
		root: absRoot,
	}

	for _, opt := range opts {
		opt(w)
	}

	return w, nil
}

type Option func(*Walker)

func WithExcludeDirs(dirs []string) Option {
	return func(w *Walker) {
		w.excludeDirs = dirs
	}
}

func WithExcludeFiles(patterns []string) Option {
	return func(w *Walker) {
		compiledPatterns := make([]glob.Glob, 0, len(patterns))
		for _, pattern := range patterns {
			if g, err := glob.Compile(pattern); err == nil {
				compiledPatterns = append(compiledPatterns, g)
			}
		}
		w.excludeFiles = compiledPatterns
	}
}

func (w *Walker) Walk() ([]string, error) {
	var files []string

	log.Printf("Starting walk from root directory: %s", w.root)

	err := filepath.Walk(w.root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("Error accessing path %s: %v", path, err)
			return nil // Continue walking despite errors
		}

		if info.IsDir() {
			if w.shouldSkipDir(path) {
				return filepath.SkipDir
			}
			// Don't return here, continue walking
			return nil
		}

		// Check if it's a file we should process
		if w.shouldSkipFile(path) {
			return nil
		}

		files = append(files, path)
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("error walking directory: %w", err)
	}

	if len(files) == 0 {
		log.Printf("No suitable files found in %s or its subdirectories", w.root)
	} else {
		log.Printf("Found %d suitable files", len(files))
	}

	return files, nil
}

func (w *Walker) shouldSkipDir(path string) bool {
	// Don't skip the root directory
	if path == w.root {
		return false
	}

	// Skip hidden directories, except for the root
	base := filepath.Base(path)
	if strings.HasPrefix(base, ".") {
		return true
	}

	// Check against excluded directories
	for _, excluded := range w.excludeDirs {
		if strings.EqualFold(base, excluded) { // Case-insensitive comparison
			log.Printf("Skipping excluded directory: %s", path)
			return true
		}
	}

	return false
}

func (w *Walker) shouldSkipFile(path string) bool {
	// Skip hidden files
	if strings.HasPrefix(filepath.Base(path), ".") {
		return true
	}

	for _, pattern := range w.excludeFiles {
		if pattern.Match(filepath.Base(path)) {
			return true
		}
	}
	return false
}