package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// cdnCmd represents the cdn command
var cdnCmd = &cobra.Command{
	Use:   "cdn",
	Short: "CDN Tools",
	Long:  `CDN Tools is a set of tools for CDN operations,`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cdn called")
	},
}

func init() {
	rootCmd.AddCommand(cdnCmd)
}
