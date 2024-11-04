Here's the updated README.md that reflects the new CLI structure:

```markdown
# codemap

`codemap` is a tool that helps annotate code files with structural information to provide better context for LLMs (Large Language Models). It adds a single line of metadata at the top of each file containing the file path, package name and language information.

## Installation

```bash
brew install krzko/tap/codemap
```

Or clone and build manually:

```bash
git clone https://github.com/krzko/codemap.git
cd codemap
make build
```

Or use docker:

```bash
docker run -v $(pwd):/app ghcr.io/krzko/codemap apply --dry-run
```

## Commands

### Apply Annotations

```bash
# Add annotations to files in current directory
codemap apply

# Add annotations with dry-run (show what would change)
codemap apply --dry-run

# Add annotations to specific directory
codemap apply -d /path/to/project

# Add annotations with verbose output
codemap apply -V
```

### Clean Annotations

```bash
# Remove annotations from files in current directory
codemap clean

# Remove annotations with dry-run
codemap clean --dry-run

# Clean specific directory
codemap clean -d /path/to/project
```

### List Files

```bash
# List files that would be processed
codemap list
# or
codemap ls

# List files in specific directory
codemap list -d /path/to/project
```

### Show Statistics

```bash
# Show annotation statistics for current directory
codemap stats

# Show stats for specific directory
codemap stats -d /path/to/project
```

### Common Options

All commands support these options:
- `-d, --dir`: Directory to process (default: current directory)
- `-t, --types`: Comma-separated list of file extensions (default: "go,py,js,jsx,ts,tsx")
- `-V, --verbose`: Enable verbose logging
- `-v, --version`: Display version information

### Example Annotation

Before:
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

After:
```go
// codemap: path=/path/to/file.go;pkg=main;lang=Go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

### Supported Languages

- Go (.go)
- Python (.py)
- JavaScript (.js, .jsx)
- TypeScript (.ts, .tsx)
- Dockerfile

### Default Exclusions

The tool automatically excludes these common directories and files:

Directories:
- .git
- .github
- .gitlab
- .vscode
- .idea
- node_modules
- vendor
- .venv
- __pycache__
- dist
- build

Files:
- Hidden files (.*) 
- *.min.js
- *.map
- *.lock
- package-lock.json
- yarn.lock
- pnpm-lock.yaml
- *.sum
- *.mod

## Examples

```bash
# Add annotations with dry-run
codemap apply --dry-run

# Process only Go files
codemap apply -t go

# Clean TypeScript files with dry-run
codemap clean -t ts,tsx --dry-run

# List all Python and JavaScript files
codemap list -t py,js

# Show stats for Go files
codemap stats -t go

# Process all files with verbose output
codemap apply -V
```
