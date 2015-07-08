// haiti-core project main.go
package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/ccpony86/haiti/core"
	"github.com/ccpony86/haiti/im"
	"github.com/ccpony86/haiti/pay"
	"github.com/ccpony86/haiti/schedule"
	"github.com/ccpony86/haiti/sso"
)

func main() {
	core.Router_register()
	sso.Router_register()
	im.Router_register()
	schedule.Router_register()
	pay.Router_register()

	fmt.Println(beego.AppConfig.String("mysqluser"))
	fmt.Println(beego.AppConfig.String("mysqlpass"))
	fmt.Println(beego.AppConfig.String("mysqlurls"))
	fmt.Println(beego.AppConfig.String("mysqldb"))

	var FilterUser = func(ctx *context.Context) {
		if v := ctx.Input.Query("access_token"); v != "" {
			//data := "ok"
			fmt.Println(v)
			//ctx.Output.Json(data, false, true)
		} else {
			data := `{"error","err"}`
			ctx.Output.Json(data, false, true)
		}
	}

	beego.InsertFilter("*", beego.BeforeRouter, FilterUser)

	beego.Run()
}
