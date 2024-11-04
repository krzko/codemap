package languages

import "strings"

type JavaScript struct{}

func (js *JavaScript) FileExtensions() []string {
	return []string{".js", ".jsx", ".ts", ".tsx"}
}

func (js *JavaScript) CommentStart() string {
	return "//"
}

func (js *JavaScript) CommentEnd() string {
	return ""
}

func (js *JavaScript) MultiLineCommentStart() string {
	return "/*"
}

func (js *JavaScript) IsSpecialComment(line string) bool {
	specialPrefixes := []string{
		"@ts-ignore",
		"@ts-nocheck",
		"@ts-check",
		"@flow",
		"// eslint-",
		"/* eslint-",
	}

	for _, prefix := range specialPrefixes {
		if strings.HasPrefix(strings.TrimSpace(line), prefix) {
			return true
		}
	}
	return false
}