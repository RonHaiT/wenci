package words

import (
	"context"

	v1 "wenci/api/words/v1"
	"wenci/internal/logic/words"
)

func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	uid, err := c.users.GetUid(ctx)
	if err != nil {
		return nil, err
	}
	err = c.words.Update(ctx, req.Id, words.UpdateInput{
		Uid:                uid,
		Word:               req.Word,
		Definition:         req.Definition,
		ExampleSentence:    req.ExampleSentence,
		ChineseTranslation: req.ChineseTranslation,
		Pronunciation:      req.Pronunciation,
		ProficiencyLevel:   req.ProficiencyLevel,
	})

	return nil, err
}
