package ls

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"syscall"

	"github.com/spf13/cobra"
)

var (
	// Do not ignore entries starting with .
	all bool
	// Use long listing format
	long bool
)

type File struct {
	ftype            string
	permission       string
	dirs             int
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

		files, err := listFiles(path)

		if err != nil {
			log.Fatal(err)
		}

		for _, f := range files {
			output := fmt.Sprintf("%s  ", f.name)

			if long {
				output = fmt.Sprintf("%v %v %v %v %v %v %v %s \n", f.ftype, f.permission, f.dirs, f.group, f.user, f.size, f.modificationDate, output)
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
		return []File{}, err
	}

	for _, d := range files {
		info, _ := d.Info()
		ftype := d.Type().String()[0]
		dirs := 1
		lastModification := info.ModTime().Format("Jan 2 15:04")
		permission := info.Mode().Perm().String()

		if d.Name()[0] == '.' && !all {
			continue
		}

		if d.Type().IsDir() {
			links, err := countLinks(filepath.Join(path, d.Name()))

			if err != nil {
				return []File{}, err
			}

			dirs = links
		}

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

		userGroup, err := user.LookupGroupId(fmt.Sprintf("%d", GID))

		if err != nil {
			return []File{}, err
		}

		owner, err := user.LookupId(fmt.Sprintf("%d", UID))

		if err != nil {
			return []File{}, err
		}

		result = append(result, File{
			ftype:            string(ftype),
			permission:       permission,
			dirs:             dirs,
			group:            userGroup.Name,
			user:             owner.Username,
			size:             int(info.Size()),
			modificationDate: lastModification,
			name:             d.Name(),
		})
	}

	return result, err
}
