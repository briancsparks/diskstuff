package cmd

/* Copyright © 2022 Brian C Sparks <briancsparks@gmail.com> -- MIT (see LICENSE file) */

import (
  "github.com/briancsparks/diskstuff/diskstuff"
  "github.com/spf13/cobra"
)

var path string
var dryrun bool
var safe bool
var verbose bool

// fixnamesCmd represents the fixnames command
var fixnamesCmd = &cobra.Command{
	Use:   "fixnames",
	Short: "A brief description of your command",

	Run: func(cmd *cobra.Command, args []string) {
    conf := diskstuff.NewConfig(dryrun)
    conf.SetSafe(safe)
    conf.SetVerbose(verbose)

    diskstuff.FixNames(path, conf)
	},
}

func init() {
	rootCmd.AddCommand(fixnamesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
  fixnamesCmd.PersistentFlags().StringVar(&path, "path", ".", "The path to scan")
  fixnamesCmd.PersistentFlags().BoolVar(&dryrun, "dryrun", false, "Show what would happen, but don't run")
  fixnamesCmd.PersistentFlags().BoolVar(&safe, "safe", false, "Do not clobber destination files")
  fixnamesCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "Show info")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fixnamesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
