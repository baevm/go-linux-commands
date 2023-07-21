//go:build !windows
// +build !windows

package ls

import (
	"fmt"
	"os"
	"syscall"
)

func countLinks(dirPath string) (int, error) {
	fi, err := os.Stat(dirPath)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	nlink := int(0)

	if sys := fi.Sys(); sys != nil {
		if stat, ok := sys.(*syscall.Stat_t); ok {
			nlink = int(stat.Nlink)
		}
	}

	return nlink, err
}
