package form

type BlogInfo struct {
	// 登录前通过用户输入获取
	UserAccount string `json:"account" mapstructure:"account"`
	AccessToken string `json:"token" mapstructure:"token"`
	BlogAddress string `json:"address" mapstructure:"address"`

	// 根据登录后的返回获取
	UserName string `json:"username" mapstructure:"username"`
}
