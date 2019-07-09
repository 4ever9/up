package main

import (
	"github.com/4ever9/up"
	"github.com/spf13/cobra"
)

var all bool

func init() {
	versionCMD.Flags().BoolVar(&all, "all", false, "show all version info")
}

var versionCMD = &cobra.Command{
	Use:   "version [flags]",
	Short: "Show version about app",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Printf("Up version: %s-%s\n", up.CurrentVersion, up.CurrentCommit)
		if all {
			cmd.Printf("App build date: %s\n", up.BuildDate)
			cmd.Printf("System version: %s\n", up.Platform)
			cmd.Printf("Golang version: %s\n", up.GoVersion)
		}
	},
}
