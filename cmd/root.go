package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "replace <old pattern> <new pattern> [files]",
	Short: "text pattern replacer for text files",
	Long: `Replace all found instances of the first argument
with the second argument. '*' is any number of characters.

If no files are given then all files under the current directory,
recursively descending into subdirectories, will be taken.`,
	Args: cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		return replace(args)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.PersistentFlags().BoolP("exact", "e", false, "Do not add asterixes at the beginning and end of patterns. Default is to add them.")

	viper.BindPFlags(rootCmd.PersistentFlags())
}
