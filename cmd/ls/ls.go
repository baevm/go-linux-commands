package ls

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/dezzerlol/go-linux-commands/internal/color"
	"github.com/spf13/cobra"
)

var (
	// Do not ignore entries starting with .
	all bool
	// Use long listing format
	long bool
)

var (
	dirKey = "d"
	libKey = "l"
)

type File struct {
	ftype            string
	permission       string
	links            int
	user             string
	group            string
	size             int
	modificationDate string
	name             string
}

func init() {
	LsCmd.Flags().BoolVarP(&long, "long", "l", false, "Use long listing format")
	LsCmd.Flags().BoolVarP(&all, "all", "a", false, "Do not ignore entries starting with .")
}

// BUG: panic when trying to read files like swapfile.sys
var LsCmd = &cobra.Command{
	Use:     "ls [OPTIONS] [FILEs]",
	Short:   "ls lists all files in directory",
	Long:    `ls lists all files in directory`,
	Version: "1.0.0",
	Args:    cobra.RangeArgs(0, 99),
	Run: func(cmd *cobra.Command, args []string) {
		path := "./"

		if len(args) > 0 {
			path = args[0]
		}

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered. Error:\n", r)
			}
		}()

		files, err := listFiles(path)

		if err != nil {
			log.Fatal(err)
		}

		for _, f := range files {
			output := fmt.Sprintf("%s  ", f.name)

			if long {
				output = fmt.Sprintf("%s %s %v %s %s %4v\t %s %s \n",
					f.ftype,
					f.permission,
					f.links,
					f.group,
					f.user,
					f.size,
					f.modificationDate,
					output)
			}

			fmt.Print(output)
		}

		fmt.Print("\n")
	},
}

func listFiles(path string) ([]File, error) {
	var result []File

	files, err := os.ReadDir(path)

	if err != nil {
		log.Println(err)
		return []File{}, err
	}

	for _, d := range files {
		info, _ := d.Info()

		linksCount := 1
		ftype := strings.ToLower(string(d.Type().String()[0]))
		lastModification := info.ModTime().Format("Jan 2 15:04")
		permission := info.Mode().Perm().String()
		fileName := d.Name()

		if fileName[0] == '.' && !all {
			continue
		}

		if ftype == dirKey {
			fileName = color.ColorStr(fileName, color.Blue)
		}

		if ftype == libKey {
			fileName = color.ColorStr(fileName, color.Cyan)
		}

		if d.Type().IsDir() {
			links, err := countLinks(filepath.Join(path, d.Name()))

			if err != nil {
				log.Println(err)
				continue
			}

			linksCount = links
		}

		owner, group, err := getFileOwners(info)

		if err != nil {
			log.Println(err)
			continue
		}

		result = append(result, File{
			ftype:            string(ftype),
			permission:       permission,
			links:            linksCount,
			group:            group.Name,
			user:             owner.Username,
			size:             int(info.Size()),
			modificationDate: lastModification,
			name:             fileName,
		})
	}

	return result, err
}

func getFileOwners(info fs.FileInfo) (*user.User, *user.Group, error) {
	var UID int
	var GID int

	if stat, ok := info.Sys().(*syscall.Stat_t); ok {
		UID = int(stat.Uid)
		GID = int(stat.Gid)
	} else {
		// we are not in linux, this won't work anyway in windows,
		// but maybe you want to log warnings
		UID = os.Getuid()
		GID = os.Getgid()
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
