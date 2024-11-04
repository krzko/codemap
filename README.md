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
docker run -v $(pwd):/app krzko/codemap
```

## Usage

Basic usage:
```bash
# Process files in current directory
codemap

# Process files in a specific directory
codemap -dir=/path/to/your/project

# Remove annotations
codemap -clean

# Enable verbose logging
codemap -verbose
```

### Command Line Options

- `-dir`: Directory to process (default: current directory)
- `-clean`: Remove existing annotations
- `-verbose`: Enable verbose logging
- `-types`: Comma-separated list of file extensions to process (default: "go,py,js,jsx,ts,tsx")
- `-version`: Print version information

### Example

Processing a file will add a single line annotation at the top:

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
# Process all supported files in current directory
codemap

# Process only Go files in current directory
codemap -types=go

# Process Python and JavaScript files in a specific directory
codemap -dir=/path/to/project -types=py,js

# Clean annotations from all files in current directory
codemap -clean

# Clean annotations from TypeScript files
codemap -clean -types=ts,tsx

# Process all supported files with verbose logging
codemap -verbose
```
