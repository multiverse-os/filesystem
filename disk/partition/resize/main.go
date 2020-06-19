package main

// TODO: test/fix on disks with non-512 byte sectors ( /sys/block/sda/queue/hw_sector_size)

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
)

var (
	dry     = flag.Bool("dry-run", false, "don't make changes")
	verbose = flag.Bool("verbose", false, "verbose output")
)

func init() {
	flag.Usage = usage
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of embiggen-disk:\n\n")
	fmt.Fprintf(os.Stderr, "# embiggen-disk [flags] <mount-point-to-enlarge>\n\n")
	flag.PrintDefaults()
	os.Exit(1)
}

func fatalf(format string, args ...interface{}) {
	log.SetFlags(0)
	log.Fatalf(format, args...)
}

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		usage()
	}
	if runtime.GOOS != "linux" {
		fatalf("embiggen-disk only runs on Linux.")
	}

	mnt := flag.Arg(0)
	e, err := getFileSystemResizer(mnt)
	if err != nil {
		fatalf("error preparing to enlarge %s: %v", mnt, err)
	}
	changes, err := Resize(e)
	if len(changes) > 0 {
		fmt.Printf("Changes made:\n")
		for _, c := range changes {
			fmt.Printf("  * %s\n", c)
		}
	} else if err == nil {
		fmt.Printf("No changes made.\n")
	}
	if err != nil {
		fatalf("error: %v", err)
	}
}

// An Resizer is anything that can enlarge something and describe its state.
// An Resizer can depend on another Resizer to run first.
type Resizer interface {
	String() string                       // "ext4 filesystem at /", "LVM PV foo"
	State() (string, error)               // "534 blocks"
	Resize() error                        // both may be non-zero
	DepResizer() (dep Resizer, err error) // can return (nil, nil) for none
}

// Resize resizes e's dependencies and then resizes e.
func Resize(e Resizer) (changes []string, err error) {
	s0, err := e.State()
	if err != nil {
		return
	}
	dep, err := e.DepResizer()
	if err != nil {
		return
	}
	if dep != nil {
		changes, err = Resize(dep)
		if err != nil {
			return
		}
	}
	err = e.Resize()
	if err != nil {
		return
	}
	s1, err := e.State()
	if err != nil {
		err = fmt.Errorf("error after successful resize of %v: %v", e, err)
		return
	}
	if s0 != s1 {
		changes = append(changes, fmt.Sprintf("%v: before: %v, after: %v", e, s0, s1))
	}
	return
}
