package filesystem

import (
	"fmt"
	"os"
	"path/filepath"
)

type Read func(path Path) (bool, error)     // Read=[Copy(Read+Create)]
type Create func(path Path) (bool, error)   // Create=[Touch,Write,Append(Read+Create)]
type Move func(from, to Path) (bool, error) // Move=[Rename]
type Delete func(path Path) (bool, error)

type Actions struct {
	Read   Read
	Create Create
	Move   Move
	Delete Delete
}

func DefaultCreate(file_type, perm os.FileMode, path Path) (bool, error) {
	switch file_type {
	case FileType:
		// TODO: Need to create paths leading to basename (file) if they do not
		// exit. Then create the file.

		// TODO: Need a way to convert path to directory
		parent_directory := filepath.Dir(path)
		if !parent_directory.IsCurrentPath() {
			os.Mkdir
			return true, nil
		} else {
			return true, nil
		}

	case DirectoryType:

		return true, nil

	default: // InvalidType
		return false, fmt.Errorf("error: invalid or unspecified type")
	}
}
