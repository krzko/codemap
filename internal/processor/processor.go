package processor

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/krzko/codemap/pkg/annotator"
	"github.com/krzko/codemap/pkg/walker"
)

type Processor struct {
	opts      Options
	annotator annotator.Annotator
	walker    *walker.Walker
}

// New creates a new Processor instance
func New(opts Options) (*Processor, error) {
	w, err := walker.New(
		opts.Directory,
		walker.WithExcludeDirs(opts.ExcludeDirs),
		walker.WithExcludeFiles(opts.ExcludeFiles),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize walker: %w", err)
	}

	return &Processor{
		opts:      opts,
		annotator: annotator.New(),
		walker:    w,
	}, nil
}

// Process handles the file processing
func (p *Processor) Process() error {
	files, err := p.walker.Walk()
	if err != nil {
		return fmt.Errorf("failed to walk directory: %w", err)
	}

	log.Printf("Found %d files to process", len(files))

	if p.opts.Clean {
		log.Printf("Running in clean mode - removing annotations")
	} else {
		log.Printf("Running in add mode - adding annotations")
	}

	if p.opts.Concurrent {
		log.Printf("Processing files concurrently with %d workers", p.opts.MaxWorkers)
		return p.processConcurrent(files)
	}

	log.Printf("Processing files sequentially")
	return p.processSequential(files)
}

func (p *Processor) processConcurrent(files []string) error {
	var wg sync.WaitGroup
	errChan := make(chan error, len(files))
	semaphore := make(chan struct{}, p.opts.MaxWorkers)

	for _, file := range files {
		wg.Add(1)
		go func(f string) {
			defer wg.Done()
			semaphore <- struct{}{}        // Acquire
			defer func() { <-semaphore }() // Release

			if err := p.processFile(f); err != nil {
				errChan <- fmt.Errorf("failed to process file %s: %w", f, err)
			}
		}(file)
	}

	wg.Wait()
	close(errChan)

	// Collect any errors
	for err := range errChan {
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *Processor) processSequential(files []string) error {
	for _, file := range files {
		if err := p.processFile(file); err != nil {
			return fmt.Errorf("failed to process file %s: %w", file, err)
		}
	}
	return nil
}

func (p *Processor) processFile(path string) error {
	// Skip unsupported files
	if !p.isSupported(path) {
		return nil
	}

	relPath, err := filepath.Rel(p.opts.Directory, path)
	if err != nil {
		relPath = path // Fallback to absolute path if relative path fails
	}

	if p.opts.Clean {
		log.Printf("Cleaning annotations from: %s", relPath)
		return p.annotator.RemoveAnnotation(path)
	}

	log.Printf("Adding annotations to: %s", relPath)
	info := annotator.FileInfo{
		Path:        path,
		Language:    p.determineLanguage(path),
		ImportPath:  p.determineImportPath(path),
		PackageName: p.determinePackageName(path),
	}

	err = p.annotator.AddAnnotation(info)
	if err != nil {
		log.Printf("Error processing %s: %v", relPath, err)
		return err
	}

	return nil
}

// Add the missing methods
func (p *Processor) determineLanguage(path string) string {
	ext := filepath.Ext(path)
	switch ext {
	case ".go":
		return "Go"
	case ".py":
		return "Python"
	case ".js", ".jsx":
		return "JavaScript"
	case ".ts", ".tsx":
		return "TypeScript"
	default:
		return "Unknown"
	}
}

func (p *Processor) determineImportPath(path string) string {
	// This is a simplistic implementation - you might want to make it more sophisticated
	// based on your project's structure
	abs, err := filepath.Abs(path)
	if err != nil {
		return path
	}

	// For Go files, try to determine the module path
	if filepath.Ext(path) == ".go" {
		// You might want to read go.mod file to get the module name
		return p.determineGoImportPath(abs)
	}

	return abs
}

func (p *Processor) determineGoImportPath(absPath string) string {
	// This is a placeholder - you should implement proper Go module detection
	// by reading go.mod file and determining the module path
	return absPath
}

func (p *Processor) determinePackageName(path string) string {
	// For Go files, read the package name from the file
	if filepath.Ext(path) == ".go" {
		return p.readGoPackageName(path)
	}

	// For other files, use the directory name as package name
	return filepath.Base(filepath.Dir(path))
}

func (p *Processor) isSupported(path string) bool {
	ext := filepath.Ext(path)
	for _, supported := range p.opts.SupportedTypes {
		if ext == supported {
			return true
		}
	}
	return false
}

func (p *Processor) readGoPackageName(path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		return "unknown"
	}

	// Simple package name extraction - you might want to make this more robust
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		if strings.HasPrefix(strings.TrimSpace(line), "package ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "package "))
		}
	}

	return "unknown"
}