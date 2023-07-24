package mkdir

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func init() {

}

var Cmd = &cobra.Command{
	Use:     "mkdir",
	Short:   "Create directory",
	Long:    `Create the DIRECTORY(ies), if they do not already exist.`,
	Version: "1.0.0",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dirName := args[0]

		err := createDirectory(dirName)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("OK.")
	},
}

func createDirectory(name string) error {
	_, err := os.Stat(name)

	if err != nil {
		if os.IsExist(err) {
			return fmt.Errorf("mkdir: cannot create directory '%s': File exists", name)
		}
	}

	err = os.Mkdir(name, 0755)

	return err
}
