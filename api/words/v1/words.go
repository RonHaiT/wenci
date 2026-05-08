package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type ProficiencyLevel uint

const (
	ProficiencyLevel1 ProficiencyLevel = iota + 1
	ProficiencyLevel2
	ProficiencyLevel3
	ProficiencyLevel4
	ProficiencyLevel5
)

type CreateReq struct {
	g.Meta             `path:"words" method:"post" sm:"创建" tags:"单词" security:"AuthToken"`
	Word               string           `json:"word" v:"required|length:1,100" dc:"单词"`
	Definition         string           `json:"definition" v:"required|length:1,300" dc:"单词定义"`
	ExampleSentence    string           `json:"example_sentence" v:"required|length:1,300" dc:"例句"`
	ChineseTranslation string           `json:"chinese_translation" v:"required|length:1,300" dc:"中文翻译"`
	Pronunciation      string           `json:"pronunciation" v:"required|length:1,100" dc:"发音"`
	ProficiencyLevel   ProficiencyLevel `json:"proficiency_level" v:"required|between:1,5" dc:"熟练度，1最低，5最高"`
}

type CreateRes struct {
}

type UpdateReq struct {
	g.Meta             `path:"words/{id}" method:"put" sm:"更新" tags:"单词" security:"AuthToken"`
	Id                 uint             `json:"id" v:"required" dc:"单词ID"`
	Word               string           `json:"word" v:"required|length:1,100" dc:"单词"`
	Definition         string           `json:"definition" v:"required|length:1,300" dc:"单词定义"`
	ExampleSentence    string           `json:"example_sentence" v:"required|length:1,300" dc:"例句"`
	ChineseTranslation string           `json:"chinese_translation" v:"required|length:1,300" dc:"中文翻译"`
	Pronunciation      string           `json:"pronunciation" v:"required|length:1,100" dc:"发音"`
	ProficiencyLevel   ProficiencyLevel `json:"proficiency_level" v:"required|between:1,5" dc:"熟练度，1最低，5最高"`
}
type UpdateRes struct{}

type List struct {
	Id               uint             `json:"id" dc:"单词ID"`
	Word             string           `json:"word" dc:"单词"`
	Definition       string           `json:"definition" dc:"单词定义"`
	ProficiencyLevel ProficiencyLevel `json:"proficiency_level" dc:"熟练度，1最低，5最高"`
}
type ListReq struct {
	g.Meta   `path:"words" method:"get" sm:"列表" tags:"单词" security:"AuthToken"`
	Word     string `json:"word" v:"length:1,100" dc:"模糊查询单词"`
	Page     int    `json:"page" v:"min:1" dc:"页码,默认1"`
	PageSize int    `json:"page_size" v:"between:1,100" dc:"每页条数,默认10"`
}

type ListRes struct {
	List  []List `json:"list" dc:"单词列表"`
	Total int    `json:"total" dc:"总条数"`
}

type DetailReq struct {
	g.Meta `path:"words/{id}" method:"get" sm:"详情" tags:"单词" security:"AuthToken"`
	Id     uint `json:"id" v:"required" dc:"单词ID"`
}
type DetailRes struct {
	Id                 uint             `json:"id" dc:"单词ID"`
	Word               string           `json:"word" dc:"单词"`
	Definition         string           `json:"definition" dc:"单词定义"`
	ExampleSentence    string           `json:"example_sentence" dc:"例句"`
	ChineseTranslation string           `json:"chinese_translation" dc:"中文翻译"`
	Pronunciation      string           `json:"pronunciation" dc:"发音"`
	ProficiencyLevel   ProficiencyLevel `json:"proficiency_level" dc:"熟练度，1最低，5最高"`
	CreatedAt          *gtime.Time      `json:"createdAt" dc:"创建时间"`
	UpdatedAt          *gtime.Time      `json:"updatedAt" dc:"更新时间"`
}
