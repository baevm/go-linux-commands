package cat

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func init() {}

var Cmd = &cobra.Command{
	Use:     "cat [FILE]",
	Short:   "Concatenate files and print on the standard output.",
	Long:    `Concatenate files and print on the standard output.`,
	Version: "1.0.0",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		res, err := printFile(name)

		if err != nil {
			log.Fatal(err)
		}

		for _, v := range res {
			fmt.Println(v)
		}

	},
}

func printFile(name string) ([]string, error) {
	file, err := os.Open(name)

	if err != nil {
		return nil, err
	}

	var res []string

	scan := bufio.NewScanner(file)

	for scan.Scan() {
		res = append(res, scan.Text())
	}

	if scan.Err() != nil {
		return nil, scan.Err()
	}

	return res, nil
}
