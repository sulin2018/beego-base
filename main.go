package main

import (
	"net/http"

	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/config/yaml" // config yaml驱动
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"

	_ "github.com/astaxie/beego/session/mysql" // session mysql驱动
	"github.com/sulin2018/beego-base/backend/models"
	"github.com/sulin2018/beego-base/backend/utils"
)

func init() {
	InitLogs()
	InitConfig()
	models.InitModels()
}

func main() {
	beego.Get("/ping", func(ctx *context.Context) {
		_ = ctx.Output.JSON(map[string]interface{}{"code": http.StatusOK, "message": "pong"}, false, false)
	})
	beego.Run()
}

func InitConfig() {
	err := beego.LoadAppConfig("yaml", "config.yaml")
	if err != nil {
		logs.Error(err)
	} else {
		logs.Info("Init config success")
	}
}

func InitLogs() {
	err := utils.MkDir("logs")
	if err != nil {
		logs.Error(err)
	}

	// 文件名行号输出
	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(3)

	// 输出日志到文件
	err = logs.SetLogger(logs.AdapterFile, `{
			"filename":"logs/beegobase.log",
			"level":6
		}`)
	if err != nil {
		logs.Error(err)
	}
	// Debug7 Info6 Warn4 Error3 由高到底，大于当前日志级别的日志不展示
	/*
		filename 保存的文件名
		maxlines 每个文件保存的最大行数，默认值 1000000
		maxsize 每个文件保存的最大尺寸，默认值是 1 << 28, //256 MB
		daily 是否按照每天 logrotate，默认是 true
		maxdays 文件最多保存多少天，默认保存 7 天
		rotate 是否开启 logrotate，默认是 true
		level 日志保存的时候的级别，默认是 Trace 级别, 即 Debug
		perm 日志文件权限
	*/

	// 同时输出到console
	_ = logs.SetLogger(logs.AdapterConsole)

	// 日志测试
	// logs.Debug("Debug")
	// logs.Info("Info")
	// logs.Warn("Warn")
	// logs.Error("Error")

	// 初始化默认日志 最好不要再使用默认log 而使用logs
	// log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	// log.SetPrefix("[ log ] ")

	// if beego.BConfig.RunMode == "prod" {
	// 	F, err := gotools.MustOpen("zping.log", "logs")
	// 	if err != nil {
	// 		logs.Error("初始化日志出错", err)
	// 	}
	// 	log.SetOutput(F)
	// }
}
