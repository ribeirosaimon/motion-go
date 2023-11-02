package domain

import (
	"fmt"

	"github.com/spf13/cobra"
)

type commandMotion struct {
	Use  string
	Help string
	Flag CommandFlagMotion
}

type CommandFlagMotion struct {
	Name         string
	ShortHand    string
	DefaultValue interface{}
	Usage        string
}

func NewCommandFlagMotion(name, short, usage string, defaultValue interface{}) *CommandFlagMotion {
	return &CommandFlagMotion{
		Name:         name,
		ShortHand:    short,
		Usage:        usage,
		DefaultValue: defaultValue,
	}
}

func NewCommandMotion(use, help string, function func(*cobra.Command, []string), flags ...*CommandFlagMotion) *cobra.Command {
	var command = &cobra.Command{
		Use:   use,
		Short: help,
		Run:   function,
	}
	for _, flag := range flags {
		command.Flags().StringP(flag.Name, flag.ShortHand, fmt.Sprintf("%v", flag.DefaultValue), flag.Usage)
	}
	return command
}
