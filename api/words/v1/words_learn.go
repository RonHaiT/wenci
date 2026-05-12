package v1

import "github.com/gogf/gf/v2/frame/g"

type RandListReq struct {
	g.Meta `path:"words/rand" method:"get" sm:"随机获取单词列表" tags:"单词" security:"AuthToken"`
	Limit  uint `json:"limit" v:"between:1,300" dc:"限制个数,默认50"`
}
type RandListRes struct {
	List []List `json:"list" dc:"单词列表"`
}

type SetLevelReq struct {
	g.Meta `path:"words/{id}/level" method:"patch" sm:"设置单词熟练度" tags:"单词" security:"AuthToken"`
	Id     uint `json:"id" v:"required" dc:"单词ID"`
	Level  uint `json:"level" v:"required|between:1,5" dc:"熟练度,1最低,5最高"`
}
type SetLevelRes struct {
}
