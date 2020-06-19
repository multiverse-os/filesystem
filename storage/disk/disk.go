package filesystem

type SuperBlock struct {
}

type INodeTable struct {
	// TODO: This poins to the datablocks,
	// Typical deletion just unlinks this from the datablocks, a real secure
	// delte requires 0'ing the datablocks or writing random data then 0'ing
	// preferably.
}

type INode struct {
}

type DataBlocks struct {
}

type File struct {
	// TODO: Each file has number of links, these ar ethe links to a datablock
	// -rw-r--r--  1 user user 1.3K Jun 19 07:42
	//             ^
	//             |  _________________________________________________
	//             '-| This is the link count, number of hard links to |
	//               | datablocks                                      |
	//               '-------------------------------------------------'
}
