package words

import (
	"context"

	v1 "wenci/api/words/v1"
	"wenci/internal/logic/words"
)

func (c *ControllerV1) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {
	uid, err := c.users.GetUid(ctx)
	if err != nil {
		return nil, err
	}
	wordList, total, err := c.words.List(ctx, words.ListInput{
		Uid:      uid,
		Word:     req.Word,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	var list []v1.List
	for _, v := range wordList {
		list = append(list, v1.List{
			Id:               v.Id,
			Word:             v.Word,
			Definition:       v.Definition,
			ProficiencyLevel: v1.ProficiencyLevel(v.ProficiencyLevel),
		})
	}
	return &v1.ListRes{
		List:  list,
		Total: total,
	}, nil
}
