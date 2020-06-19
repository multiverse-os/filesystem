package filesystem

// Non-Atomic
// fd=open("file", O_TRUNC); write(fd, data); close(fd);
// Atomic
// fd=open("file.new"); write(fd, data); close(fd); rename("file.new", "file");

func New(file_type FileType, path string) interface{} {
	switch file_type {
	case FileType:
		return *File{
			Path: Path{path},
		}
	case DirectoryType:
		return *Directory{
			Path:        Path{path},
			Directories: []*Directory{},
			Files:       []*File{},
		}
	default:
		return nil
	}
}

////////////////////////////////////////////////////////////////////////////////
// NOTE: Only add logic that is actively being used, we will have a different
// Filesystem package for full/complete implementation of all linux POSIX
// functionality. This is just what is needed for the minimum/basic application
// interactions with the filesystem.

type FilesystemActions struct {
	File      Actions
	Directory Actions
}

type RootDirectory *Directory

type RootFilesystem struct {
	Root          *RootDirectory
	BeforeActions FilesystemActions
	Actions       FilesystemActions
	AfterActions  FilesystemActions
}

var Filesystem = RootFilesystem{
	Root: *RootDirectory{},
	Actions: FilesystemActions{
		File: Actions{
			Create: DefaultCreate(Path),
		},
	},
}
