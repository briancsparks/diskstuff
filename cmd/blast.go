/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
  "github.com/briancsparks/diskstuff/diskstuff"

  "github.com/spf13/cobra"
)

// blastCmd represents the blast command
var blastCmd = &cobra.Command{
	Use:   "blast",
	Short: "A brief description of your command",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("blast called")
    diskstuff.Blast("")
	},
}

func init() {
	rootCmd.AddCommand(blastCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// blastCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// blastCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
