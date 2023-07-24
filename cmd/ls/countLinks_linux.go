//go:build !windows
// +build !windows

package ls

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/user"
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

func getFileOwners(info fs.FileInfo) (*user.User, *user.Group, error) {
	var UID int
	var GID int

	if stat, ok := info.Sys().(*syscall.Stat_t); ok {
		UID = int(stat.Uid)
		GID = int(stat.Gid)
	}

	group, err := user.LookupGroupId(fmt.Sprintf("%d", GID))

	if err != nil {
		log.Println(err)
		return nil, nil, err
	}

	owner, err := user.LookupId(fmt.Sprintf("%d", UID))

	if err != nil {
		log.Println(err)
		return nil, nil, err
	}

	return owner, group, nil
}
