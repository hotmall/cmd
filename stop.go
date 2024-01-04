/*
Copyright © 2023 Jack Ju <jack.ju@itiky.com>
This file is part of CLI application foo.
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop one running service",
	Run: func(cmd *cobra.Command, args []string) {
		exit(LogPath())
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
