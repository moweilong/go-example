package consts

// 状态码
const (
	StatusALL     int = -1 // 全部状态
	StatusEnabled int = 1  // 启用
	StatusDisable int = 2  // 禁用
	StatusDelete  int = 3  // 已删除
)

// RequestEncryptKey
// 请求加密密钥用于敏感数据加密，16位字符，前后端需保持一致
// 安全起见，生产环境运行时请注意修改
var RequestEncryptKey = []byte("f080a463654b2279")
