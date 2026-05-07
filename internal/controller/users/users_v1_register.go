package users

import (
	"context"

	v1 "wenci/api/users/v1"
	"wenci/internal/logic/users"
)

func (c *ControllerV1) Register(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error) {
	err = c.users.Register(ctx, users.RegisterInput{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	})
	return nil, err
}
