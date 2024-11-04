package annotator

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/krzko/codemap/internal/languages"
)

const (
	annotationPattern = "codemap: path="
)

// DefaultAnnotator implements the Annotator interface
type DefaultAnnotator struct {
	languages map[string]languages.Language
}

func New() Annotator {
	return &DefaultAnnotator{
		languages: map[string]languages.Language{
			".go":         &languages.GoLang{},
			".py":         &languages.Python{},
			".js":         &languages.JavaScript{},
			".jsx":        &languages.JavaScript{},
			".ts":         &languages.JavaScript{},
			".tsx":        &languages.JavaScript{},
			".dockerfile": &languages.Dockerfile{},
			"":            &languages.Dockerfile{},
		},
	}
}

func (a *DefaultAnnotator) AddAnnotation(info FileInfo) error {
	lang, ok := a.languages[filepath.Ext(info.Path)]
	if !ok {
		return fmt.Errorf("unsupported file type: %s", info.Path)
	}

	// Read the file content
	content, err := os.ReadFile(info.Path)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %v", info.Path, err)
	}

	// Check if annotation already exists
	if a.hasAnnotationWithLang(string(content), lang) {
		relPath, err := filepath.Rel(".", info.Path)
		if err != nil {
			relPath = info.Path
		}
		log.Printf("Skipping file (already annotated): %s", relPath)
		return nil
	}

	// Create annotation and write
	annotation := a.createAnnotation(lang, info)
	f, err := os.OpenFile(info.Path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %v", info.Path, err)
	}
	defer f.Close()

	if _, err := f.WriteString(annotation + string(content)); err != nil {
		return fmt.Errorf("failed to write file %s: %v", info.Path, err)
	}

	relPath, err := filepath.Rel(".", info.Path)
	if err != nil {
		relPath = info.Path
	}
	log.Printf("Added annotations to: %s", relPath)

	return nil
}

func (a *DefaultAnnotator) RemoveAnnotation(path string) error {
	lang, ok := a.languages[filepath.Ext(path)]
	if !ok {
		return fmt.Errorf("unsupported file type: %s", path)
	}

	// Read the file
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	if !a.hasAnnotationWithLang(string(content), lang) {
		relPath, err := filepath.Rel(".", path)
		if err != nil {
			relPath = path
		}
		log.Printf("Skipping file (no annotations): %s", relPath)
		return nil
	}

	// Process the file line by line
	var lines []string
	scanner := bufio.NewScanner(strings.NewReader(string(content)))
	isFirstLine := true

	for scanner.Scan() {
		line := scanner.Text()
		// Skip only if it's the first line and contains our annotation
		if isFirstLine && strings.Contains(line, annotationPattern) {
			isFirstLine = false
			continue
		}
		lines = append(lines, line)
		isFirstLine = false
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	// Write the file back
	if err := os.WriteFile(path, []byte(strings.Join(lines, "\n")), 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %v", path, err)
	}

	relPath, err := filepath.Rel(".", path)
	if err != nil {
		relPath = path
	}
	log.Printf("Removed annotations from: %s", relPath)

	return nil
}

func (a *DefaultAnnotator) createAnnotation(lang languages.Language, info FileInfo) string {
	commentStart := lang.CommentStart()
	return fmt.Sprintf("%s codemap: path=%s;pkg=%s;lang=%s\n",
		commentStart,
		info.Path,
		info.PackageName,
		info.Language)
}

// HasAnnotation checks if a file has a codemap annotation
func (a *DefaultAnnotator) HasAnnotation(content string) bool {
	scanner := bufio.NewScanner(strings.NewReader(content))
	// Only check the first line
	if !scanner.Scan() {
		return false
	}
	firstLine := scanner.Text()
	return strings.Contains(firstLine, annotationPattern)
}

// hasAnnotationWithLang is an internal helper that checks for language-specific annotation
func (a *DefaultAnnotator) hasAnnotationWithLang(content string, lang languages.Language) bool {
	scanner := bufio.NewScanner(strings.NewReader(content))
	// Only check the first line
	if !scanner.Scan() {
		return false
	}
	firstLine := scanner.Text()
	commentStart := lang.CommentStart()
	return strings.Contains(firstLine, commentStart+" "+annotationPattern)
}