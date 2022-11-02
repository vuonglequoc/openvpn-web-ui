package models

import (
	"os"

	"github.com/vuonglequoc/go-openvpn/server/config"
	"github.com/beego/beego"
	"github.com/beego/beego/orm"
	passlib "gopkg.in/hlandau/passlib.v1"
)

var GlobalCfg Settings

func init() {
	initDB()
	createDefaultUsers()
	createDefaultSettings()
	createDefaultOVConfig()
}

func initDB() {
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	dbSource := "file:" + beego.AppConfig.String("dbPath")

	err := orm.RegisterDataBase("default", "sqlite3", dbSource)
	if err != nil {
		panic(err)
	}
	orm.Debug = true
	orm.RegisterModel(
		new(User),
		new(Settings),
		new(OVConfig),
	)

	// Database alias.
	name := "default"
	// Drop table and re-create.
	force := false
	// Print log.
	verbose := true

	err = orm.RunSyncdb(name, force, verbose)
	if err != nil {
		beego.Error(err)
		return
	}
}

func createDefaultUsers() {
	hash, err := passlib.Hash(os.Getenv("OPENVPN_ADMIN_PASSWORD"))
	if err != nil {
		beego.Error("Unable to hash password", err)
	}
	user := User{
		Id:       1,
		Login:    os.Getenv("OPENVPN_ADMIN_USERNAME"),
		Name:     "Administrator",
		Email:    "root@localhost",
		Password: hash,
	}
	o := orm.NewOrm()
	if created, _, err := o.ReadOrCreate(&user, "Name"); err == nil {
		if created {
			beego.Info("Default admin account created")
		} else {
			beego.Debug(user)
		}
	}

}

func createDefaultSettings() {
	s := Settings{
		Profile:       "default",
		MIAddress:     "openvpn:2080",
		MINetwork:     "tcp",
		ServerAddress: "127.0.0.1",
		OVConfigPath:  "/etc/openvpn/",
		CAConfigPath:  "/etc/ca_server/",
	}
	o := orm.NewOrm()
	if created, _, err := o.ReadOrCreate(&s, "Profile"); err == nil {
		GlobalCfg = s

		if created {
			beego.Info("New settings profile created")
		} else {
			beego.Debug(s)
		}
	} else {
		beego.Error(err)
	}
}

func createDefaultOVConfig() {
	c := OVConfig{
		Profile: "default",
		Config: config.Config{
			Port:                1194,
			Proto:               "udp",
			Ca:                  "/etc/openvpn/pki/ca.crt",
			Cert:                "/etc/openvpn/pki/"+os.Getenv("SERVER_NAME")+".crt",
			Key:                 "/etc/openvpn/pki/private/"+os.Getenv("SERVER_NAME")+".key",
			Dh:                  "/etc/openvpn/pki/dh2048.pem",
			Server:              "10.8.0.0 255.255.255.0",
			IfconfigPoolPersist: "/etc/openvpn/log/ipp.txt",
			Keepalive:           "10 120",
			TaKey:               "/etc/openvpn/pki/ta.key",
			Cipher:              "AES-256-CBC",
			Auth:                "SHA256",
			MaxClients:          100,
			Management:          "0.0.0.0 2080",
		},
	}
	o := orm.NewOrm()
	if created, _, err := o.ReadOrCreate(&c, "Profile"); err == nil {
		if created {
			beego.Info("New settings profile created")
		} else {
			beego.Debug(c)
		}
		path := GlobalCfg.OVConfigPath + "server.conf"
		if _, err = os.Stat(path); os.IsNotExist(err) {
			destPath := GlobalCfg.OVConfigPath + "server.conf"
			if err = config.SaveToFile("conf/openvpn-server-config.tpl",
				c.Config, destPath); err != nil {
				beego.Error(err)
			}
		}
	} else {
		beego.Error(err)
	}
}
