package touch

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
)

func init() {

}

var Cmd = &cobra.Command{
	Use:     "touch [OPTION] [FILE]",
	Short:   "Create empty file",
	Long:    `Create empty file`,
	Version: "1.0.0",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]

		err := createFile(filename)

		if err != nil {
			log.Fatal(err)
		}
	},
}

func createFile(filename string) error {
	isExist, err := Exists(filename)

	if err != nil {
		return fmt.Errorf("touch: cannot create file '%s': ", err)
	}

	if isExist {
		err := os.Chtimes(filename, time.Now().Local(), time.Now().Local())
		return err
	}

	_, err = os.Create(filename)

	return err
}

func Exists(name string) (bool, error) {
	_, err := os.Stat(name)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}
