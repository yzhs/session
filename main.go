package main

import (
	"os"
	"os/exec"

	"flag"

	"gopkg.in/mattes/go-expand-tilde.v1"
)

const (
	BASE_DIR = "~/.sessions"
)

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || !os.IsNotExist(err)
}

func main() {
	flag.Parse()
	if flag.NArg() > 0 {
		base_dir, err := tilde.Expand(BASE_DIR)
		if err != nil {
			panic(err)
		}
		for _, arg := range flag.Args() {
			// Actually reload a session
			fullPath := base_dir + "/" + arg
			if fileExists(fullPath) {
				cmd := exec.Command("zsh", fullPath)
				err := cmd.Run()
				if err != nil {
					panic(err)
				}
			} else {
				println("File not found: " + fullPath)
			}
		}
	} else {
		flag.Usage()
	}
}
