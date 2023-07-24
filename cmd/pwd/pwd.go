package pwd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func init() {

}

var Cmd = &cobra.Command{
	Use:     "pwd",
	Short:   "Print name of current/working directory",
	Long:    `Print name of current/working directory`,
	Version: "1.0.0",
	Run: func(cmd *cobra.Command, args []string) {
		directory, err := getCurrentDir()

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(directory)
	},
}

func getCurrentDir() (string, error) {
	path, err := os.Getwd()

	return path, err
}
