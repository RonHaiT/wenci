package main

import (
	"errors"
	_ "wenci/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"

	"wenci/internal/cmd"
)

func main() {
	var err error

	// 全局设置i18n

	g.I18n().SetLanguage("cn")

	// 检查数据库是否能连接
	err = connDb()
	if err != nil {
		panic(err)
	}

	cmd.Main.Run(gctx.GetInitCtx())
}

// connDb 检查数据库连接是否正常
func connDb() error {
	err := g.DB().PingMaster()
	if err != nil {
		return errors.New("连接到数据库失败")
	} else {
		g.Log().Info(gctx.GetInitCtx(), "数据库连接成功")
		return nil
	}
}
