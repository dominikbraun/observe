package main

import (
	"github.com/spf13/cobra"
	"os"
)

var version string = "UNSPECIFIED"

func main() {
	var ctx context

	root := &cobra.Command{
		Use:     "observe <url> <settings path>",
		Short:   `Observe a website and get an e-mail if something changes.`,
		Version: version,
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	website := &cobra.Command{
		Use:   "website <url> <settings path>",
		Short: `Observe a website`,
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			settings, err := readSettings(args[1])
			if err != nil {
				return err
			}

			ctx.settings = settings
			ctx.url = args[0]

			return observeWebsite(&ctx, os.Stdout)
		},
	}

	root.PersistentFlags().UintVarP(&ctx.interval, "interval", "i", 1, `interval for checks`)
	root.PersistentFlags().BoolVarP(&ctx.quitOnChange, "quit-on-change", "q", false, `stop observing after a change`)
	root.AddCommand(website)

	if err := root.Execute(); err != nil {
		panic(err)
	}
}
