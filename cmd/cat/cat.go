package cat

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	// Prefix each found line with number
	lineNumber bool
	// Number nonempty output lines. overrides -n
	nonBlank bool
)

func init() {
	Cmd.Flags().BoolVarP(&lineNumber, "number", "n", false, "Prefix each found line with number")
	Cmd.Flags().BoolVarP(&nonBlank, "number-nonblank", "b", false, "Number nonempty output lines. overrides -n")
}

var Cmd = &cobra.Command{
	Use:     "cat [FILE]",
	Short:   "Concatenate files and print on the standard output.",
	Long:    `Concatenate files and print on the standard output.`,
	Version: "1.0.0",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		res, err := getFileData(name)

		if err != nil {
			log.Fatal(err)
		}

		for i, v := range res {
			output := v

			if lineNumber && !nonBlank {
				output = fmt.Sprintf("%d\t%s", i+1, output)
			}

			if nonBlank {
				if v != "" {
					output = fmt.Sprintf("%d\t%s", i+1, output)
				}
			}

			fmt.Println(output)
		}

	},
}

func getFileData(name string) ([]string, error) {
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
