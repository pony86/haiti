// pkg1 project pkg1.go
package core

import (
	"fmt"
	"github.com/astaxie/beego"
)

func Info() {
	fmt.Println("core module")
}

func Run() {
	fmt.Println("core run")

	beego.RESTRouter("/object", &ObjectController{})
	beego.Run()
}
