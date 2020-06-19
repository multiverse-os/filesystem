package main

import (
	"fmt"

	filesystem "../../filesystem"
)

func main() {
	fmt.Println("nested directories example")
	fmt.Println("==========================")
	fmt.Println("A simple example that illustrates how a developer can create")
	fmt.Println("a local data path for an application with several directories")
	fmt.Println("and subdirectories.")

	// NOTE: Should treat this as a virtual filesystem (VFS), letting us do
	// lookups against this filesystem without requiring us knowing the base; like
	// a chroot in a way.

	// TODO: Should we have a virtualsystem be an object that doesn't create the
	// files? Should we not create the objects until we flush or write? Then we
	// cna have a function take care of writing that can be changed.
	directory := filesystem.New(filesystem.Directory, "/home/user/")

	dotlocal := directory.AddDirectory(".local")

	share := dotlocal.AddDirectory("share")
	app_data := share.AddDirectory("app")
}
