package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type InfoReq struct {
	g.Meta `path:"account/info" method:"get" sm:"获取账户信息" tags:"用户" security:"AuthToken"`
}
type InfoRes struct {
	Username  string      `json:"username" dc:"用户名"`
	Email     string      `json:"email" dc:"邮箱"`
	CreatedAt *gtime.Time `json:"created" dc:"创建时间"`
	UpdatedAt *gtime.Time `json:"updated" dc:"更新时间"`
}
