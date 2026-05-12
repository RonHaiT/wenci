package words

import (
	"context"
	v1 "wenci/api/words/v1"
	"wenci/internal/dao"
	"wenci/internal/model/entity"

	"github.com/gogf/gf/v2/errors/gerror"
)

func (w *Words) Rand(ctx context.Context, uid, limit uint) ([]entity.Words, error) {
	if limit <= 0 {
		limit = 50
	}
	var (
		err  error
		cls  = dao.Words.Columns()
		orm  = dao.Words.Ctx(ctx)
		list []entity.Words
	)
	if uid > 0 {
		orm = orm.Where(cls.Uid, uid)
	}
	err = orm.Limit(int(limit)).OrderRandom().Scan(&list)
	return list, err
}

func (w *Words) SetLevel(ctx context.Context, uid, id uint, level v1.ProficiencyLevel) error {
	if level < v1.ProficiencyLevel1 || level > v1.ProficiencyLevel5 {
		return gerror.New("熟练度值不合法")
	}

	var (
		cls = dao.Words.Columns()
		orm = dao.Words.Ctx(ctx)
	)
	if uid > 0 {
		orm = orm.Where(cls.Uid, uid)
	}
	_, err := orm.Data(cls.ProficiencyLevel, uint(level)).Where(cls.Id, id).Where(cls.Uid, uid).Update()

	return err
}
