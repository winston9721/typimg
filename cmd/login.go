package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"runtime"
	"typoraImg/config"
	"typoraImg/consts"
	"typoraImg/form"
	"typoraImg/service"
	"typoraImg/utils"
)

func getUserInput() form.BlogInfo {
	pt := utils.NewPrint(runtime.GOOS)
	pt.ColorPrintln(
		fmt.Sprintf(
			"注意事项：\n %s\n %s\n",
			"1.打开【https://i.cnblogs.com/settings页面->其他设置】，根据以下提示输入配置中的相应项",
			"2.进行第1步的同时，请勾选其他设置里的 允许 MetaWeblog 博客客户端访问，随后点击保存，非常重要！！！",
		),
		consts.ColorRed,
	)

	pt.ColorPrintln("请根据以下提示输入，按回车结束：", consts.ColorYellow)
	reader := bufio.NewReader(os.Stdin)

	pt.ColorPrint(fmt.Sprintf("%-12s", "MetaWeblog登录名："), consts.ColorYellow)
	account, _ := reader.ReadString('\n')

	pt.ColorPrint(fmt.Sprintf("%-12s", "MetaWeblog访问令牌："), consts.ColorYellow)
	password, _ := reader.ReadString('\n')

	pt.ColorPrint(fmt.Sprintf("%-12s", "MetaWeblog访问地址："), consts.ColorYellow)
	rpcAddr, _ := reader.ReadString('\n')

	var blogInfo = form.BlogInfo{
		//UserAccount: strings.Replace(account, "\n", "", -1),
		UserAccount: utils.Input.Format(account),
		//AccessToken: strings.Replace(password, "\n", "", -1),
		AccessToken: utils.Input.Format(password),
		BlogAddress: utils.Input.Format(rpcAddr),
	}
	return blogInfo
}

func initCommandFunc() {
	pt := utils.NewPrint(runtime.GOOS)

	// 获取用户输入的登录信息
	blogInfo := getUserInput()

	// 模拟登录，根据返回的数据回填blogInfo信息
	err := service.LoginBlog(&blogInfo)
	utils.Error.ExitIf(err)

	// 更新或者创建配置文件
	err = config.UpCreateConfig(blogInfo)
	utils.Error.ExitIf(err)

	// 登录成功提示语
	pt.ColorPrintln(
		fmt.Sprintf("\n登录成功，欢迎你，%s!\n", blogInfo.UserName),
		consts.ColorGreen,
	)

	// 提示命令
	abPath, err := utils.Path.GetExecutableAbPath()
	utils.Error.ExitIf(err)

	pt.ColorPrintln(
		"请将以下命令复制到【Typora->设置->图像->图像上传服务(选中自定义命令)->命令输入框】中",
		consts.ColorGreen,
	)
	pt.ColorPrintln(
		fmt.Sprintf("%s %s\n", abPath, "upload"),
		consts.ColorGreen,
	)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "初始化博客园登录信息",
	Long:  `初始化上传图片需要的博客园登录信息，仅初始化一次即可永久使用`,
	Run: func(cmd *cobra.Command, args []string) {
		initCommandFunc()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
