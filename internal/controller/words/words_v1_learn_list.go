package words

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"wenci/api/words/v1"
)

func (c *ControllerV1) LearnList(ctx context.Context, req *v1.LearnListReq) (res *v1.LearnListRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
