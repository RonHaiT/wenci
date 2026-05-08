package words

import (
	"context"
	v1 "wenci/api/words/v1"
	"wenci/internal/dao"
	"wenci/internal/model/do"

	"github.com/gogf/gf/v2/errors/gerror"
)

type Words struct{}

func New() *Words {
	return &Words{}
}

type CreateInput struct {
	Uid                uint   `json:"uid" v:"required" dc:"用户ID"`
	Word               string `json:"word" v:"required|length:1,100" dc:"单词"`
	Definition         string `json:"definition" v:"required|length:1,300" dc:"单词定义"`
	ExampleSentence    string `json:"example_sentence" v:"required|length:1,300" dc:"例句"`
	ChineseTranslation string `json:"chinese_translation" v:"required|length:1,300" dc:"中文翻译"`
	Pronunciation      string `json:"pronunciation" v:"required|length:1,100" dc:"发音"`
	ProficiencyLevel   v1.ProficiencyLevel
}

func (w *Words) Create(ctx context.Context, in CreateInput) error {
	var cls = dao.Words.Columns()
	count, err := dao.Words.Ctx(ctx).Where(cls.Word, in.Word).Where(cls.Word, in.Word).Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return gerror.New("单词已存在")
	}
	_, err = dao.Words.Ctx(ctx).Data(do.Words{
		Uid:                in.Uid,
		Word:               in.Word,
		Definition:         in.Definition,
		ExampleSentence:    in.ExampleSentence,
		ChineseTranslation: in.ChineseTranslation,
		Pronunciation:      in.Pronunciation,
		ProficiencyLevel:   in.ProficiencyLevel,
	}).Insert()
	if err != nil {
		return err
	}
	return nil
}

type UpdateInput struct {
	Uid                uint   `json:"id" v:"required" dc:"单词ID"`
	Word               string `json:"word" v:"required|length:1,100" dc:"单词"`
	Definition         string `json:"definition" v:"required|length:1,300" dc:"单词定义"`
	ExampleSentence    string `json:"example_sentence" v:"required|length:1,300" dc:"例句"`
	ChineseTranslation string `json:"chinese_translation" v:"required|length:1,300" dc:"中文翻译"`
	Pronunciation      string `json:"pronunciation" v:"required|length:1,100" dc:"发音"`
	ProficiencyLevel   v1.ProficiencyLevel
}

func (w *Words) Update(ctx context.Context, id uint, in UpdateInput) error {
	var cls = dao.Words.Columns()
	count, err := dao.Words.Ctx(ctx).Where(cls.Uid, in.Uid).Where(cls.Word, in.Word).WhereNot(cls.Id, id).Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return gerror.New("单词已存在")
	}
	_, err = dao.Words.Ctx(ctx).Data(do.Words{
		Word:               in.Word,
		Definition:         in.Definition,
		ExampleSentence:    in.ExampleSentence,
		ChineseTranslation: in.ChineseTranslation,
		Pronunciation:      in.Pronunciation,
		ProficiencyLevel:   in.ProficiencyLevel,
	}).Where(cls.Id, id).Where(cls.Uid, in.Uid).Update()
	if err != nil {
		return err
	}
	return nil
}
