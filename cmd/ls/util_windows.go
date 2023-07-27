//go:build windows
// +build windows

package ls

import (
	"io/fs"
	"os/user"
)

func countLinks(path string, filename string) (int, error) {
	return 0, nil
}

func getFileOwners(info fs.FileInfo) (*user.User, *user.Group, error) {
	return &user.User{}, &user.Group{}, nil
}
