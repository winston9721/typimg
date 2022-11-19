package cmd

import (
	"github.com/spf13/cobra"
	"typoraImg/utils"
)

var rootCmd = &cobra.Command{
	Use:   "typimg",
	Short: "上传图片到博客园",
	Long:  `上传图片到博客园`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//    Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		utils.Error.ExitIf(err)
	}
}
