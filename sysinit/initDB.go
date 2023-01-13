package sysinit

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"     // 引入orm的包
	_ "github.com/go-sql-driver/mysql" // 引入MySQL包
)

func InitDB() {
	// 连接名称 别名
	dbAlias := beego.AppConfig.String("db_alias")
	// 数据库名称
	dbName := beego.AppConfig.String("db_name")
	// 数据库连接用户名
	dbUser := beego.AppConfig.String("db_user")
	// 数据库连接密码
	dbPwd := beego.AppConfig.String("db_pwd")
	// 数据库IP（域名）
	dbHost := beego.AppConfig.String("db_host")
	// 数据库端口
	dbPort := beego.AppConfig.String("db_port")
	// 编码格式
	dbCharset := beego.AppConfig.String("db_charset")

	// 注册数据库
	orm.RegisterDataBase(dbAlias, "mysql", dbUser+":"+dbPwd+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?"+dbCharset, 30)

	// 如果是开发模式， 则显示命令信息
	isDev := (beego.AppConfig.String("runmode") == "dev")
	// 非强制模式下自动建表
	err := orm.RunSyncdb("default", false, isDev)
	if err != nil {
		logs.Informational("[orm] Create table err : ", err)
	}
	if isDev {
		orm.Debug = isDev
	}
}
