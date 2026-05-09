package v1

import "github.com/gogf/gf/v2/frame/g"

type LearnListReq struct {
	g.Meta `path:"words/rand" method:"get" sm:"随机获取单词列表" tags:"单词" security:"AuthToken"`
	Limit  uint `json:"limit" v:"between:1,300" dc:"限制个数,默认50"`
}
type LearnListRes struct {
	List []List `json:"list" dc:"单词列表"`
}
