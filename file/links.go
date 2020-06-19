package filesystem

type LinkType int

const (
	Soft Link = iota
	Hard
)

// NOTE Soft links have different INodes
//      Hard links have same INodes

type Link struct {
	Metadata Metadata
	Link     Path
}
