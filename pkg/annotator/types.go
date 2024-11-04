package annotator

type FileInfo struct {
	Path        string
	Language    string
	ImportPath  string
	PackageName string
}

type Annotator interface {
	// AddAnnotation adds file structure information to the file
	AddAnnotation(info FileInfo) error
	// RemoveAnnotation removes existing annotation from the file
	RemoveAnnotation(path string) error
}