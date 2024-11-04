package languages

import "strings"

type Python struct{}

func (p *Python) FileExtensions() []string {
	return []string{".py"}
}

func (p *Python) CommentStart() string {
	return "#"
}

func (p *Python) CommentEnd() string {
	return ""
}

func (p *Python) MultiLineCommentStart() string {
	return `"""`
}

func (p *Python) IsSpecialComment(line string) bool {
	specialPrefixes := []string{
		"# type:",
		"# noqa:",
		"# pylint:",
		"# pragma:",
	}

	for _, prefix := range specialPrefixes {
		if strings.HasPrefix(strings.TrimSpace(line), prefix) {
			return true
		}
	}
	return false
}