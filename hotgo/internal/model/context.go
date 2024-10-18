package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Identity 通用身份模型
type Identity struct {
	Id       int64       `json:"id"              description:"用户ID"`
	Pid      int64       `json:"pid"             description:"上级ID"`
	DeptId   int64       `json:"deptId"          description:"部门ID"`
	DeptType string      `json:"deptType"        description:"部门类型"`
	RoleId   int64       `json:"roleId"          description:"角色ID"`
	RoleKey  string      `json:"roleKey"         description:"角色唯一标识符"`
	Username string      `json:"username"        description:"用户名"`
	RealName string      `json:"realName"        description:"姓名"`
	Avatar   string      `json:"avatar"          description:"头像"`
	Email    string      `json:"email"           description:"邮箱"`
	Mobile   string      `json:"mobile"          description:"手机号码"`
	App      string      `json:"app"             description:"登录应用"`
	LoginAt  *gtime.Time `json:"loginAt"         description:"登录时间"`
}
