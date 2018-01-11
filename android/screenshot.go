package main

import (
	"os/exec"
)

func Screenshot(filename string) {
	_, err := exec.Command("/system/bin/screencap", "-p", filename).Output()
	if err != nil {
		panic("screenshot failed")
	}
}
