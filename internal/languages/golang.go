package languages

import "strings"

type GoLang struct{}

func (g *GoLang) FileExtensions() []string {
	return []string{".go"}
}

func (g *GoLang) CommentStart() string {
	return "//"
}

func (g *GoLang) CommentEnd() string {
	return ""
}

func (g *GoLang) MultiLineCommentStart() string {
	return "/*"
}

func (g *GoLang) IsSpecialComment(line string) bool {
	specialPrefixes := []string{
		"//go:generate",
		"//go:build",
		"//nolint",
		"// +build",
	}

	for _, prefix := range specialPrefixes {
		if strings.HasPrefix(strings.TrimSpace(line), prefix) {
			return true
		}
	}
	return false
}