package routers

import (
	"github.com/beego/beego/v2/server/web/context/param"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {

    beego.GlobalControllerRouter["github.com/vuonglequoc/openvpn-web-ui/controllers:APISessionController"] = append(beego.GlobalControllerRouter["github.com/vuonglequoc/openvpn-web-ui/controllers:APISessionController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/vuonglequoc/openvpn-web-ui/controllers:APISessionController"] = append(beego.GlobalControllerRouter["github.com/vuonglequoc/openvpn-web-ui/controllers:APISessionController"],
        beego.ControllerComments{
            Method: "Kill",
            Router: `/`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/vuonglequoc/openvpn-web-ui/controllers:APISignalController"] = append(beego.GlobalControllerRouter["github.com/vuonglequoc/openvpn-web-ui/controllers:APISignalController"],
        beego.ControllerComments{
            Method: "Send",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/vuonglequoc/openvpn-web-ui/controllers:APISysloadController"] = append(beego.GlobalControllerRouter["github.com/vuonglequoc/openvpn-web-ui/controllers:APISysloadController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/vuonglequoc/openvpn-web-ui/controllers:CertificatesController"] = append(beego.GlobalControllerRouter["github.com/vuonglequoc/openvpn-web-ui/controllers:CertificatesController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/certificates`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/vuonglequoc/openvpn-web-ui/controllers:CertificatesController"] = append(beego.GlobalControllerRouter["github.com/vuonglequoc/openvpn-web-ui/controllers:CertificatesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/certificates`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/vuonglequoc/openvpn-web-ui/controllers:CertificatesController"] = append(beego.GlobalControllerRouter["github.com/vuonglequoc/openvpn-web-ui/controllers:CertificatesController"],
        beego.ControllerComments{
            Method: "Download",
            Router: `/certificates/download/:key`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/vuonglequoc/openvpn-web-ui/controllers:CertificatesController"] = append(beego.GlobalControllerRouter["github.com/vuonglequoc/openvpn-web-ui/controllers:CertificatesController"],
        beego.ControllerComments{
            Method: "Renew",
            Router: `/certificates/renew/:key`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/vuonglequoc/openvpn-web-ui/controllers:CertificatesController"] = append(beego.GlobalControllerRouter["github.com/vuonglequoc/openvpn-web-ui/controllers:CertificatesController"],
        beego.ControllerComments{
            Method: "Revoke",
            Router: `/certificates/revoke/:key`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
