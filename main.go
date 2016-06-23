package main

import (
	"os"
	"os/exec"

	"flag"
)

const (
	BASE_DIR = "/home/joghurt/.sessions"
)

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || !os.IsNotExist(err)
}



func main() {
	flag.Parse()
	if flag.NArg() > 0 {
		for _, arg := range flag.Args() {
			// Actually reload a session
			fullPath := BASE_DIR + "/" + arg
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

