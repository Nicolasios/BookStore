package main

import (
	_ "BookStore/routers"
	"BookStore/sysinit"

	c "github.com/TreyBastian/colourize"

	"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.Debug = true
	sysinit.InitDB()

}

func main() {
	logs.SetLogger(logs.AdapterMultiFile, `{"filename":"logs/mylog.log","separate":["error","mergency","info"]}`) //日志分类输出到各个文件

	logs.Info(c.Colourize("begin beego framework", c.Black, c.Whitebg, c.Bold))
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(
		&cors.Options{
			AllowAllOrigins:  true,
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
			ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
			AllowCredentials: true}))
	beego.Run()
}
