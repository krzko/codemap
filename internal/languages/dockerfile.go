package languages

type Dockerfile struct{}

func (d *Dockerfile) FileExtensions() []string {
    return []string{".dockerfile", ""}  // Empty string for files named exactly "Dockerfile"
}

func (d *Dockerfile) CommentStart() string {
    return "#"
}

func (d *Dockerfile) CommentEnd() string {
    return ""
}

func (d *Dockerfile) MultiLineCommentStart() string {
    return "#"
}

func (d *Dockerfile) IsSpecialComment(line string) bool {
    return false
}