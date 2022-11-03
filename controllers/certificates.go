package controllers

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/vuonglequoc/go-openvpn/client/config"
	"github.com/vuonglequoc/openvpn-web-ui/lib"
	"github.com/vuonglequoc/openvpn-web-ui/models"
	"github.com/beego/beego"
	"github.com/beego/beego/validation"
)

type CertificatesController struct {
	BaseController
}

func (c *CertificatesController) NestPrepare() {
	if !c.IsLogin {
		c.Ctx.Redirect(302, c.LoginPath())
		return
	}
	c.Data["breadcrumbs"] = &BreadCrumbs{
		Title: "Certificates",
	}
}

// @router /certificates [get]
func (c *CertificatesController) Get() {
	c.TplName = "certificates.html"
	c.showCerts()
}

func (c *CertificatesController) showCerts() {
	path := models.GlobalCfg.CAConfigPath + "pki/index.txt"
	certs, err := lib.ReadCerts(path)
	if err != nil {
		beego.Error(err)
	}
	lib.Dump(certs)

	// The first Certs maybe belong to the VPN Server with CN name not server
	if certs[0].Details.CommonName == os.Getenv("SERVER_NAME") {
		c.Data["certificates"] = certs[1:]
	} else {
		c.Data["certificates"] = &certs
	}
}

// @router /certificates [post]
func (c *CertificatesController) Post() {
	c.TplName = "certificates.html"
	flash := beego.NewFlash()

	cParams := lib.NewCertParams{}
	if err := c.ParseForm(&cParams); err != nil {
		beego.Error(err)
		flash.Error(err.Error())
		flash.Store(&c.Controller)
	} else {
		if vMap := validateCertParams(cParams); vMap != nil {
			c.Data["validation"] = vMap
		} else {
			if err := lib.CreateCertificate(cParams); err != nil {
				beego.Error(err)
				flash.Error(err.Error())
				flash.Store(&c.Controller)
			}
		}
	}

	c.showCerts()
}

func validateCertParams(cert lib.NewCertParams) map[string]map[string]string {
	valid := validation.Validation{}
	b, err := valid.Valid(&cert)
	if err != nil {
		beego.Error(err)
		return nil
	}
	if !b {
		return lib.CreateValidationMap(valid)
	}
	return nil
}

// @router /certificates/revoke/:key [get]
func (c *CertificatesController) Revoke() {
	name := c.GetString(":key")
	lib.RevokeCertificate(name)
	c.Redirect(c.URLFor("CertificatesController.Get"), 302)
	return
}

// @router /certificates/renew/:key [get]
func (c *CertificatesController) Renew() {
	name := c.GetString(":key")
	lib.RenewCertificate(name)
	c.Redirect(c.URLFor("CertificatesController.Get"), 302)
	return
}

// @router /certificates/download/:key [get]
func (c *CertificatesController) Download() {
	name := c.GetString(":key")
	filename := fmt.Sprintf("%s.zip", name)

	c.Ctx.Output.Header("Content-Type", "application/zip")
	c.Ctx.Output.Header("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))

	zw := zip.NewWriter(c.Controller.Ctx.ResponseWriter)

	ovpnPath := models.GlobalCfg.OVConfigPath
	keysPath := ovpnPath + "client-configs/keys/"
	keysPathOvpn := ovpnPath + "client-configs/files/"

	if cfgPath, err := saveClientConfig(name); err == nil {
		addFileToZip(zw, cfgPath)
	}
	addFileToZip(zw, ovpnPath+"pki/ca.crt")
	addFileToZip(zw, keysPath+"client_"+name+".crt")
	addFileToZip(zw, ovpnPath+"pki/private/client_"+name+".key")
	addFileToZip(zw, ovpnPath+"pki/ta.key")

	// Make OVPN
	cmd := exec.Command(
		"/bin/sh",
		"-c",
		fmt.Sprintf(
			"cat " + keysPath+"client_"+name+".conf " +
			"<(echo -e '<ca>') " +
			ovpnPath+"pki/ca.crt " +
			"<(echo -e '</ca>\n<cert>') " +
			keysPath+"client_"+name+".crt " +
			"<(echo -e '</cert>\n<key>') " +
			ovpnPath+"pki/private/client_"+name+".key " +
			"<(echo -e '</key>\n<tls-auth>') " +
			ovpnPath+"pki/ta.key " +
			"<(echo -e '</tls-auth>') " +
			"> "+keysPathOvpn+"client_"+name+".ovpn"))
	cmd.Dir = ovpnPath + "client-configs"
	output, err := cmd.CombinedOutput()
	if err != nil {
		beego.Debug(string(output))
		beego.Error(err)
	}

	addFileToZip(zw, keysPathOvpn+"client_"+name+".ovpn")

	if err := zw.Close(); err != nil {
		beego.Error(err)
	}
}

func saveClientConfig(name string) (string, error) {
	cfg := config.New()
	serverConfig := models.OVConfig{Profile: "default"}
	serverConfig.Read("Profile")

	cfg.Proto = serverConfig.Proto
	cfg.ServerAddress = models.GlobalCfg.ServerAddress
	cfg.Port = serverConfig.Port

	cfg.Cert = "client_" + name + ".crt"
	cfg.Key = "client_" + name + ".key"

	cfg.Cipher = serverConfig.Cipher
	cfg.Auth = serverConfig.Auth

	cfg.ExtraClientOptions = serverConfig.ExtraClientOptions

	destPath := models.GlobalCfg.OVConfigPath + "client-configs/keys/client_" + name + ".conf"
	if err := config.SaveToFile("conf/openvpn-client-config.tpl",
		cfg, destPath); err != nil {
		beego.Error(err)
		return "", err
	}

	return destPath, nil
}

func addFileToZip(zw *zip.Writer, path string) error {
	header := &zip.FileHeader{
		Name:         filepath.Base(path),
		Method:       zip.Store,
		ModifiedTime: uint16(time.Now().UnixNano()),
		ModifiedDate: uint16(time.Now().UnixNano()),
	}
	fi, err := os.Open(path)
	if err != nil {
		beego.Error(err)
		return err
	}

	fw, err := zw.CreateHeader(header)
	if err != nil {
		beego.Error(err)
		return err
	}

	if _, err = io.Copy(fw, fi); err != nil {
		beego.Error(err)
		return err
	}

	return fi.Close()
}
