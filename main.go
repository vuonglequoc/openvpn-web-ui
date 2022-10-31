package main

import (
	"github.com/vuonglequoc/openvpn-web-ui/lib"
	_ "github.com/vuonglequoc/openvpn-web-ui/routers"
	"github.com/beego/beego"
)

func main() {
	lib.AddFuncMaps()
	beego.Run()
}
