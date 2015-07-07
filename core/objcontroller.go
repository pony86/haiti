package core

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

type ObjectController struct {
	beego.Controller
}

func (o *ObjectController) Post() {

}

func (o *ObjectController) Get() {
	req := httplib.Get("http://beego.me/")
	str, err := req.String()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)

	o.Data["json"] = str
	o.ServeJson()
}

func (o *ObjectController) Put() {

}

func (o *ObjectController) Delete() {

}
