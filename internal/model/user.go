package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type IdReq struct {
	Id int64 `json:"id" v:"required|min:1#id不能为空"`
}

type IdsReq struct {
	Ids []int64 `json:"ids"`
}

type UserSignInReq struct {
	Username string `p:"mobile" v:"required#用户名不能为空"`
	Password string `p:"password" v:"required#密码不能为空"`
	Nickname string `p:"nickname"`
}

type UserSignInRes struct {
	Token      string      `json:"token,omitempty"`
	Id         uint        `json:"id"         `           // 主键
	Username   string      `json:"username,omitempty"   ` // 用户名
	Nickname   string      `json:"nickname,omitempty"   ` // 昵称
	CreateTime *gtime.Time `json:"createTime,omitempty" ` // 创建时间
	UpdateTime *gtime.Time `json:"updateTime,omitempty" ` // 更新时间
	DeleteTime *gtime.Time `json:"deleteTime,omitempty" ` // 删除时间
}
