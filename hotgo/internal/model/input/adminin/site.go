package adminin

// LoginModel 统一登录响应
type LoginModel struct {
	Id       int64  `json:"id"              dc:"用户ID"`
	Username string `json:"username"        dc:"用户名"`
	Token    string `json:"token"           dc:"登录token"`
	Expires  int64  `json:"expires"         dc:"登录有效期"`
}

// AccountLoginInp 账号登录
type AccountLoginInp struct {
	Username string `json:"username" v:"required#用户名不能为空" dc:"用户名"`
	Password string `json:"password" v:"required#密码不能为空" dc:"密码"`
	// Cid      string `json:"cid"  dc:"验证码ID"`
	// Code     string `json:"code" dc:"验证码"`
	// IsLock   bool   `json:"isLock"  dc:"是否为锁屏状态"`
}
