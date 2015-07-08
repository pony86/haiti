// pkg1 project pkg1.go
package core

import (
	"fmt"
	"github.com/astaxie/beego"
)

func Router_register() {
	fmt.Println("core router register")

	beego.RESTRouter("/object", &ObjectController{})
}
