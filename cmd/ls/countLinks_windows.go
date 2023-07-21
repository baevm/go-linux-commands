//go:build windows
// +build windows

package ls

import (
	"os"
)

func countLinks(dirPath string) (int, error) {
	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	list, err := f.Readdirnames(-1)
	f.Close()

	return len(list), err
}
