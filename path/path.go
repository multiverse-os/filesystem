package path

import (
	"path/filepath"
	"strings"
)

type Path string

var Cusor = Path{"/"}

// TODO: Add validation
// TODO: Add segmentation
// TODO: Add base, path, absolute, ~ substituon
// TODO: Add type determination

type Segment string

type Base string
type Extension string

// TODO: For these, we want to replace before any operations
// TODO: Add ability to add more shortcuts like these
const (
	Root      Path = "/"
	Back      Path = "-"
	BackTwice Path = "--"
	Home      Path = "~"
	Current   Path = "."
	Parent    Path = ".."
)

func (self Path) IsCurrentDirectory() bool { return (self == Current) }
func (self Path) IsParentDirectory() bool  { return (self == Parent) }
func (self Path) IsRootDirectory() bool    { return (self == Root) }

// TODO: Need to create paths leading to basename (file)
// if they do not exit. Then create the file.
func (self Path) Segments() []Segment {
	return []Segment{
		strings.Split("/", string(self)),
	}
}

func (self Path) ParentPath() bool {
	//if 2 =< len(parent_directory) && byte(parent_directory[:1]) == byte("/") {}
	// NOTE Dir will return "." if there is none.
	return filepath.Dir(string(self))
}

// Abs returns name if name is an absolute path. If name is a relative
// path then an absolute path is constructed by using cwd as the current
// working directory.
func Abs(cwd, name string) string {
	if filepath.IsAbs(name) {
		return name
	}
	return filepath.Join(cwd, name)
}
