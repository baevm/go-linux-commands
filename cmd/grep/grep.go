package grep

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/dezzerlol/go-linux-commands/internal/color"
	"github.com/spf13/cobra"
)

type Result struct {
	line    string
	lineNum int
	path    string
}

var (
	hiddenPath = '.'
)

var (
	// Prefix each found line with number
	lineNumbers bool
	// Output line count
	lineCount bool
	// Recursive search in directory
	recursive bool
)

func init() {
	GrepCmd.Flags().BoolVarP(&lineNumbers, "line-number", "n", false, "Prefix each found line with number")
	GrepCmd.Flags().BoolVarP(&lineCount, "line-count", "c", false, "Output line count")
	GrepCmd.Flags().BoolVarP(&recursive, "recursive", "r", false, "Recursive search in directory")
}

var GrepCmd = &cobra.Command{
	Use:     "grep [PATTERN] [FILE]",
	Short:   "Grep searches for pattern in given file",
	Long:    `Grep searches for pattern in given file`,
	Args:    cobra.MinimumNArgs(2),
	Version: "1.0.0",
	Run: func(cmd *cobra.Command, args []string) {
		pattern := args[0]
		pathToFile := args[1]

		var res []Result
		var err error

		if recursive {
			res, err = searchDirectory(pattern, pathToFile)
		} else {
			res, err = searchInFile(pattern, pathToFile)
		}

		if err != nil {
			log.Fatal(err)
		}

		// Print line count
		if lineCount {
			fmt.Println(len(res))
			return
		}

		// Print matched lines
		for _, v := range res {
			output := v.line

			// Add line numbers where pattern is matched
			if lineNumbers {
				lineNum := color.ColorStr(v.lineNum, color.Cyan)
				output = fmt.Sprintf("%v: %s", lineNum, output)
			}

			// Add file path
			if recursive {
				path := color.ColorStr(v.path, color.Purple)
				output = fmt.Sprintf("%v:%s", path, output)
			}

			fmt.Println(output)
		}

	},
}

func searchDirectory(pattern, path string) ([]Result, error) {
	var result []Result
	var recursiveSearch func(pattern, path string) error

	recursiveSearch = func(pattern, path string) error {
		files, err := os.ReadDir(path)

		if err != nil {
			return err
		}

		for _, f := range files {
			if f.Name()[0] == byte(hiddenPath) {
				continue
			}

			if f.IsDir() {
				recursiveSearch(pattern, filepath.Join(path, f.Name()))
			} else {
				res, err := searchInFile(pattern, filepath.Join(path, f.Name()))

				if err != nil {
					return err
				}

				result = append(result, res...)
			}
		}

		return nil
	}

	err := recursiveSearch(pattern, path)

	return result, err
}

func searchInFile(pattern string, pathToFile string) ([]Result, error) {
	// Open file with given path
	file, err := os.Open(pathToFile)

	if err != nil {
		return []Result{}, err
	}

	defer file.Close()

	result := []Result{}

	if err != nil {
		return []Result{}, err
	}

	pattern = strings.ToLower(pattern)
	lineNum := 0

	// Scan file line by line
	// and check if line contains given pattern
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lineNum += 1
		line := scanner.Text()

		if strings.Contains(line, pattern) {
			coloredStr := colorizePattern(line, pattern)

			res := Result{
				line:    coloredStr,
				lineNum: lineNum,
				path:    file.Name(),
			}

			result = append(result, res)
		}
	}

	if err := scanner.Err(); err != nil {
		return []Result{}, err
	}

	return result, nil
}

func colorizePattern(line, pattern string) string {
	line = strings.TrimSpace(line)

	wordIdx := strings.Index(line, pattern)
	wordLen := len(pattern)

	endIdx := wordIdx + wordLen

	// cut first part of string
	// add color to given pattern
	// combine with end of string
	coloredStr := line[:wordIdx] + color.ColorStr(line[wordIdx:endIdx], color.Green) + line[endIdx:]

	return coloredStr
}
