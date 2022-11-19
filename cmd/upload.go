package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"typoraImg/config"
	"typoraImg/service"
	"typoraImg/utils"
)

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "上传图片服务",
	Long:  `上传图片服务 -> typimg upload img_path img_path ...`,
	Run: func(cmd *cobra.Command, args []string) {
		// 加载配置
		err := config.LoadConfig()
		utils.Error.ExitIf(err)

		// 上传
		for _, path := range args {
			if onlinePath, err := service.UploadMediaContent(path); err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println(onlinePath)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)
}
