package main

import (
	"github.com/spf13/cobra"
)

var version string = "UNSPECIFIED"

func main() {
	root := &cobra.Command{
		Use:     "observe",
		Version: version,
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	if err := root.Execute(); err != nil {
		panic(err)
	}
}
