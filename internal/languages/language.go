package languages

type Language interface {
	// FileExtensions returns the file extensions this language handles
	FileExtensions() []string
	// CommentStart returns the string that starts a single-line comment
	CommentStart() string
	// CommentEnd returns the string that ends a multi-line comment (if applicable)
	CommentEnd() string
	// MultiLineCommentStart returns the string that starts a multi-line comment
	MultiLineCommentStart() string
	// IsSpecialComment returns true if this is a special comment that shouldn't be removed
	IsSpecialComment(line string) bool
}