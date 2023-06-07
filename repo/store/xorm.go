package store

import (
	"github.com/LeeZXin/zsf/logger"
	"github.com/LeeZXin/zsf/property"
	"github.com/LeeZXin/zsf/xormlog"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var (
	Orm *xorm.Engine
)

// 初始化xorm
func init() {
	engine, err := xorm.NewEngine("mysql", property.GetString("xorm.dataSourceName"))
	if err != nil {
		logger.Logger.Panic(err)
	}
	engine.SetLogger(xormlog.XormReportLogger)
	Orm = engine
}
