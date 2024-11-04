package annotator

type FileInfo struct {
	Path        string
	Language    string
	ImportPath  string
	PackageName string
}

// Annotator interface defines the methods for file annotation handling
type Annotator interface {
	// AddAnnotation adds file structure information to the file
	AddAnnotation(info FileInfo) error
	// RemoveAnnotation removes existing annotation from the file
	RemoveAnnotation(path string) error
	// HasAnnotation checks if a file has a codemap annotation
	HasAnnotation(content string) bool
}