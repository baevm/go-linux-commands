package rmdir

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func init() {

}

var Cmd = &cobra.Command{
	Use:     "rmdir",
	Short:   "Delete directory",
	Long:    `Delete the DIRECTORY.`,
	Version: "1.0.0",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dirName := args[0]

		err := deleteDirectory(dirName)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("OK.")
	},
}

func deleteDirectory(name string) error {
	file, err := os.Open(name)

	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("mkdir: cannot create directory '%s': File doesnt exist", name)
		}

		return err
	}

	defer file.Close()

	info, err := file.Stat()

	if err != nil {
		return err
	}

	if !info.IsDir() {
		return fmt.Errorf("rmdir: failed to remove '%s': Not a directory", name)
	}

	_, err = file.Readdirnames(1)

	if err == nil {
		return fmt.Errorf("rmdir: failed to remove '%s': Directory not empty", name)
	}

	err = os.Remove(name)

	return err
}
