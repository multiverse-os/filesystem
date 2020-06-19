package metadata

import (
	"time"
)

type INode int

// Using ls -li we can see the INode ID
// TODO: Add ability to parse each element in this line
// 3282205 -rw-r--r-- 1 user user 1298 Jun 19 07:42 directory.go

type Timestamps struct {
	CreatedAt      time.Time // creation time
	LastAccessedAt time.Time // access time
	LastModifiedAt time.Time // modification time
}

type Metadata struct {
	Permissions *Permissions
	Timestamps  Timestamps
	Size        int // in kb
}
