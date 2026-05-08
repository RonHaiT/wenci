package words

import (
	"context"

	v1 "wenci/api/words/v1"
	"wenci/internal/logic/words"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	uid, err := c.users.GetUid(ctx)
	if err != nil {
		return nil, err
	}
	err = c.words.Create(ctx, words.CreateInput{
		Uid:                uid,
		Word:               req.Word,
		Definition:         req.Definition,
		ExampleSentence:    req.ExampleSentence,
		ChineseTranslation: req.ChineseTranslation,
		Pronunciation:      req.Pronunciation,
		ProficiencyLevel:   req.ProficiencyLevel,
	})
	if err != nil {
		return nil, err
	}
	return &v1.CreateRes{}, nil
}
