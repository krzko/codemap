package processor

type Options struct {
	// Directory to process
	Directory string
	// Clean mode removes annotations
	Clean bool
	// Recursive determines if we should process subdirectories
	Recursive bool
	// ExcludeDirs lists directories to skip
	ExcludeDirs []string
	// ExcludeFiles lists file patterns to skip
	ExcludeFiles []string
	// Concurrent determines if we should process files concurrently
	Concurrent bool
	// MaxWorkers limits the number of concurrent workers (0 = unlimited)
	MaxWorkers int
	// SupportedTypes lists the file extensions to process
	SupportedTypes []string
	// Verbose enables detailed logging
	Verbose bool
}

func DefaultOptions() Options {
	return Options{
		Directory: ".",
		Clean:     false,
		Recursive: true,
		ExcludeDirs: []string{
			".git",
			".github",
			".gitlab",
			".vscode",
			".idea",
			"node_modules",
			"vendor",
			".venv",
			"__pycache__",
			"dist",
			"build",
		},
		ExcludeFiles: []string{
			".*",
			"*.min.js",
			"*.map",
			"*.lock",
			"package-lock.json",
			"yarn.lock",
			"pnpm-lock.yaml",
			"*.sum",
			"*.mod",
		},
		Concurrent: true,
		MaxWorkers: 4,
		SupportedTypes: []string{
			".go",
			".py",
			".js",
			".jsx",
			".ts",
			".tsx",
			".dockerfile",
			"",
		},
		Verbose: false,
	}
}