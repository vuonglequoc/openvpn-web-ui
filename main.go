package main

import (
	"github.com/vuonglequoc/openvpn-web-ui/lib"
	_ "github.com/vuonglequoc/openvpn-web-ui/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	lib.AddFuncMaps()
	beego.Run()
}
