package main

import (
	"fmt"
	"os"

	"github.com/ribeirosaimon/motion-go/cli/commands"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "app"}

	var motionCommands = []*cobra.Command{
		commands.UpAppCommand("motion", "Up Full Motion App"),
	}

	for _, cmd := range motionCommands {
		rootCmd.AddCommand(cmd)
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
