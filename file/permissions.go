package file

import (
	"os"
)

const (
	R_OK = 0x4
	W_OK = 0x2
	X_OK = 0x1
)

// `int` value, avoiding string comparison
type UID int
type GUID int

type Group struct {
	Name string
	ID   GUID
}

type User struct {
	Name string
	ID   UID
}

type Ownership struct {
	User  *User
	Group *Group
}

type Permissions struct {
	Permission os.FileMode
	Ownership  *Ownership
}
