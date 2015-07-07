// haiti-core project main.go
package main

import (
	"flag"
	"fmt"
	//"github.com/ccpony86/haiti/core"
)

func main() {
	service_type := flag.String("service", "core", "service type")
	flag.Parse()

	fmt.Println(*service_type)

	//core.Xxx()
}
