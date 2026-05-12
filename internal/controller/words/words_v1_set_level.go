package words

import (
	"context"

	v1 "wenci/api/words/v1"
)

func (c *ControllerV1) SetLevel(ctx context.Context, req *v1.SetLevelReq) (res *v1.SetLevelRes, err error) {
	uid, err := c.users.GetUid(ctx)
	if err != nil {
		return nil, err
	}
	err = c.words.SetLevel(ctx, uid, req.Id, v1.ProficiencyLevel(req.Level))

	return nil, err
}
