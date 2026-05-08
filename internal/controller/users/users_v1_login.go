package users

import (
	"context"

	v1 "wenci/api/users/v1"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	token, err := c.users.Login(ctx, req.Username, req.Password)
	if err != nil {
		return nil, err
	}
	res = &v1.LoginRes{
		Token: token,
	}
	return res, nil
}
