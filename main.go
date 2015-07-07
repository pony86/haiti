// haiti-core project main.go
package main

import (
	"flag"
	"fmt"
	"github.com/ccpony86/haiti/core"
	"github.com/ccpony86/haiti/im"
	"github.com/ccpony86/haiti/pay"
	"github.com/ccpony86/haiti/schedule"
	"github.com/ccpony86/haiti/sso"
)

func main() {
	service_type := flag.String("service", "core", "service type")
	flag.Parse()

	fmt.Println(*service_type)

	switch *service_type {
	case "core":
		core.Run()
	case "sso":
		sso.Run()
	case "im":
		im.Run()
	case "schedule":
		schedule.Run()
	case "pay":
		pay.Run()
	}
}
