package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

type lvResizer string // /dev/mapper/debianvg-root

func (r lvResizer) String() string { return fmt.Sprintf("LVM LV %s", string(r)) }

type lvState struct {
	dev        string // 0th element in lvdisplay -c
	vg         string // 1
	numSectors int64  // 6
}

func (r lvResizer) state() (s lvState, err error) {
	s.dev = string(r)
	// # lvdisplay -c /dev/mapper/debvg-root
	//   /dev/debvg/root:debvg:3:1:-1:1:8434778112:1029636:-1:0:-1:254:0
	outb, err := exec.Command("lvdisplay", "-c", s.dev).Output()
	if err != nil {
		return s, fmt.Errorf("running lvdisplay -c %s: %v", s.dev, execErrDetail(err))
	}
	f := strings.Split(strings.TrimSpace(string(outb)), ":")
	if len(f) < 13 {
		return s, fmt.Errorf("too few expected fields in lvdisplay -c %s output: %q", s.dev, outb)
	}
	s.vg = f[1]
	s.numSectors, err = strconv.ParseInt(f[6], 10, 64)
	if err != nil {
		return s, fmt.Errorf("bogus field at index 6 in lvdisplay -c %s output: %q: %v", s.dev, outb, err)
	}
	return s, nil
}

func (r lvResizer) DepResizer() (Resizer, error) {
	lvs, err := r.state()
	if err != nil {
		return nil, err
	}

	out, err := exec.Command("pvdisplay", "-c").Output()
	if err != nil {
		return nil, fmt.Errorf("running pvdisplay -c: %v", execErrDetail(err))
	}
	bs := bufio.NewScanner(bytes.NewReader(out))
	for bs.Scan() {
		f := strings.Split(strings.TrimSpace(bs.Text()), ":")
		if len(f) < 2 || f[1] != lvs.vg {
			continue
		}
		dev := f[0]
		// TODO: support LVs with more than one PV. But that's
		// not a problem I have with cloudy things. So skip
		// for now. Probably change the DepResizer method to
		// return []Resizer.
		return pvResizer(dev), nil
	}
	return nil, nil
}

func (r lvResizer) State() (string, error) {
	lvs, err := r.state()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("sectors=%d", lvs.numSectors), nil
}

func (r lvResizer) Resize() error {
	lvDev := string(r)
	if *dry {
		fmt.Printf("[dry-run] would've run lvextend -l +100%%FREE %s", lvDev)
		return nil
	}
	_, err := exec.Command("lvextend", "-l", "+100%FREE", lvDev).Output()
	if err != nil {
		ee, ok := err.(*exec.ExitError)
		if ok && strings.Contains(string(ee.Stderr), "matches existing size") {
			return nil
		}
		var extraMsg string
		if ok && len(ee.Stderr) > 0 {
			extraMsg = fmt.Sprintf("; stderr=%s", ee.Stderr)
		}
		return fmt.Errorf("lvextend on %s: %v%s", lvDev, err, extraMsg)
	}
	return nil
}

type pvResizer string // "/dev/sda3" or potentially a whole disk e.g. "/dev/sdb"

func (r pvResizer) String() string { return fmt.Sprintf("LVM PV %s", string(r)) }

func (r pvResizer) State() (string, error) {
	dev := string(r)
	out, err := exec.Command("pvdisplay", "-c", dev).Output()
	if err != nil {
		return "", errors.New(execErrDetail(err))
	}
	f := strings.Split(strings.TrimSpace(string(out)), ":")
	if len(f) < 3 {
		return "", fmt.Errorf("bogus pvdisplay -c %s output: %q", dev, out)
	}
	return fmt.Sprintf("sectors=%v", f[2]), nil
}

func (r pvResizer) Resize() error {
	dev := string(r)
	if *dry {
		fmt.Printf("[dry-run] would've run pvresize %v", dev)
		return nil
	}
	out, err := exec.Command("pvresize", dev).CombinedOutput()
	if err != nil {
		return fmt.Errorf("pvresize %s: %v, %s", dev, err, out)
	}
	return nil
}

func (r pvResizer) DepResizer() (Resizer, error) {
	dev := string(r)
	if devEndsInNumber(dev) {
		return partitionResizer(dev), nil
	}
	return nil, nil
}
