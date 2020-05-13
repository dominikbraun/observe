// Package main provides the observe application.
package main

import (
	"github.com/spf13/cobra"
	"log"
	"os"
)

// version is used for general version information. It can be set when
// building the Go binary using `-ldflags "-X main.version=1.2.3"`.
var version string = "UNSPECIFIED"

// main is the entry point for the binary. It sets up all cobra commands
// and runs the root command.
func main() {
	var ctx context

	// root is the `observe` root command which does itself nothing but
	// provide all sub-commands as well as persisted flags like `-i`.
	root := &cobra.Command{
		Use:     "observe",
		Version: version,
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	// website implements the `observe website` command.
	website := &cobra.Command{
		Use:  "website <url> <settings path>",
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			settings, err := readSettings(args[1])
			if err != nil {
				return err
			}

			ctx.settings = settings

			return observeWebsite(&ctx, args[0], os.Stdout)
		},
	}

	root.PersistentFlags().UintVarP(&ctx.interval, "interval", "i", 1, `The interval for checks`)
	root.PersistentFlags().BoolVarP(&ctx.quitOnChange, "quit-on-change", "q", false, `Stop observing after a change`)
	root.AddCommand(website)

	if err := root.Execute(); err != nil {
		log.Fatal(err)
	}
}
