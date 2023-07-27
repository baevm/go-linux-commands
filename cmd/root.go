package cmd

import (
	"log"

	"github.com/dezzerlol/go-linux-commands/cmd/cat"
	"github.com/dezzerlol/go-linux-commands/cmd/grep"
	"github.com/dezzerlol/go-linux-commands/cmd/ls"
	"github.com/dezzerlol/go-linux-commands/cmd/mkdir"
	"github.com/dezzerlol/go-linux-commands/cmd/pwd"
	"github.com/dezzerlol/go-linux-commands/cmd/rmdir"
	"github.com/dezzerlol/go-linux-commands/cmd/touch"
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
	RootCmd.AddCommand(mkdir.Cmd)
	RootCmd.AddCommand(rmdir.Cmd)
	RootCmd.AddCommand(cat.Cmd)
	RootCmd.AddCommand(touch.Cmd)

	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
