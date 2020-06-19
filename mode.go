package filesystem

import (
	"os"
)

// NOTE [ This is very interesting ]

const (
	// The single letters are the abbreviations
	// used by the String method's formatting.
	ModeDir os.FileMode = 1 << (32 - 1 - iota) // d: is a directory
	//ModeAppend                                        // a: append-only
	//ModeExclusive                                     // l: exclusive use
	//ModeTemporary                                     // T: temporary file; Plan 9 only
	ModeSymlink   // L: symbolic link
	ModeDevice    // D: device file
	ModeNamedPipe // p: named pipe (FIFO)
	ModeSocket    // S: Unix domain socket
	//ModeSetuid                                        // u: setuid
	//ModeSetgid                                        // g: setgid
	ModeCharDevice // c: Unix character device, when ModeDevice is set
	//ModeSticky                                        // t: sticky
	ModeIrregular // ?: non-regular file; nothing else is known about this file

	// Mask for the type bits. For regular files, none will be set.
	ModeType = ModeDir | ModeSymlink | ModeNamedPipe | ModeSocket | ModeDevice | ModeCharDevice | ModeIrregular

	ModePerm os.FileMode = 0777 // Unix permission bits
)

//type FileType int
//
//const (
//	Binary FileType = iota
//	Image
//	Video
//	Audio
//	Text
//	Source
//	// ...
//)
//
//type LinkType int
//
//const (
//	Hard LinkType = iota
//	Soft
//)

// POSIX
//type Type int
//
//const (
//	Invalid       Type = iota
//	FileType           // -
//	DirectoryType      // d
//	//Device                // c = char, b = block
//	//Socket                // s
//	//Pipe                  // p
//	//Link                  // l
//)
