// pkg1 project pkg1.go
package core

import (
	"github.com/astaxie/beego"
)

func Xxx() {
	beego.RESTRouter("/object", &ObjectController{})
	beego.Run()
}
