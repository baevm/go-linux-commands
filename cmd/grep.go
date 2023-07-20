package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

type Result struct {
	data []string
}

var (
	colorGreen = "\033[32m"
	colorReset = "\033[0m"
)

func init() {
	rootCmd.AddCommand(grepCmd)
}

var grepCmd = &cobra.Command{
	Use:     "grep [PATTERN] [FILE]",
	Short:   "Grep searches for pattern in given file",
	Long:    `Grep searches for pattern in given file`,
	Args:    cobra.MinimumNArgs(2),
	Version: "1.0.0",
	Run: func(cmd *cobra.Command, args []string) {
		pattern := args[0]
		pathToFile := args[1]

		res, err := searchInFile(pattern, pathToFile)

		if err != nil {
			log.Fatal(err)
		}

		for _, v := range res.data {
			fmt.Println(v)
		}

	},
}

func searchInFile(pattern string, pathToFile string) (Result, error) {
	// Open file with given path
	file, err := os.Open(pathToFile)

	if err != nil {
		return Result{}, err
	}

	defer file.Close()

	result := Result{}

	pattern = strings.ToLower(pattern)

	// Scan file line by line
	// and check if line contains given pattern
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, pattern) {
			coloredStr := colorizePattern(line, pattern)

			result.data = append(result.data, coloredStr)
		}
	}

	if err := scanner.Err(); err != nil {
		return Result{}, err
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
	coloredStr := line[:wordIdx] + string(colorGreen) + line[wordIdx:endIdx] + string(colorReset) + line[endIdx:]

	return coloredStr
}
