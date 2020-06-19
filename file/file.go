package file

type File struct {
	Directory *Directory
	Path      Path
	Data      []byte
}
