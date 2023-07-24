package cmd

import (
	"log"
	"os"

	"github.com/dezzerlol/go-linux-commands/cmd/grep"
	"github.com/dezzerlol/go-linux-commands/cmd/ls"
	"github.com/dezzerlol/go-linux-commands/cmd/pwd"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "glc",
	Short: "glc is a linux commands written in golang",
	Long:  `glc is a linux commands written in golang`,
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	RootCmd.AddCommand(ls.Cmd)
	RootCmd.AddCommand(grep.Cmd)
	RootCmd.AddCommand(pwd.Cmd)

	if err := RootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
