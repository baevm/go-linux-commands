package cmd

import (
	"log"
	"os"

	"github.com/dezzerlol/go-linux-commands/cmd/grep"
	"github.com/dezzerlol/go-linux-commands/cmd/ls"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "glc",
	Short: "glc is a linux commands written in golang",
	Long:  `glc is a linux commands written in golang`,
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	RootCmd.AddCommand(ls.LsCmd)
	RootCmd.AddCommand(grep.GrepCmd)

	if err := RootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
