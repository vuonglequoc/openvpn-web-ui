package models

import (
	"os"

	"github.com/vuonglequoc/go-openvpn/server/config"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"

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
	dbSource, err := beego.AppConfig.String("dbPath")
	if err != nil {
		logs.Error(err)
	}
	dbSource = "file:" + dbSource

	err = orm.RegisterDataBase("default", "sqlite3", dbSource)
	if err != nil {
		logs.Error(err)
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
		logs.Error(err)
		return
	}
}

func createDefaultUsers() {
	hash, err := passlib.Hash(os.Getenv("OPENVPN_ADMIN_PASSWORD"))
	if err != nil {
		logs.Error("Unable to hash password", err)
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
			logs.Info("Default admin account created")
		} else {
			logs.Debug(user)
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
			logs.Info("New settings profile created")
		} else {
			logs.Debug(s)
		}
	} else {
		logs.Error(err)
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
			DNSServerOne:        "8.8.8.8",
			DNSServerTwo:        "8.8.4.4",
			Keepalive:           "10 120",
			TaKey:               "/etc/openvpn/pki/ta.key",
			Cipher:              "AES-256-CBC",
			Auth:                "SHA256",
			MaxClients:          100,
			Management:          "0.0.0.0 2080",
			ExtraServerOptions:  "",
			ExtraClientOptions:  "",
		},
	}
	o := orm.NewOrm()
	if created, _, err := o.ReadOrCreate(&c, "Profile"); err == nil {
		if created {
			logs.Info("New settings profile created")
		} else {
			logs.Debug(c)
		}
		path := GlobalCfg.OVConfigPath + "server.conf"
		if _, err = os.Stat(path); os.IsNotExist(err) {
			destPath := GlobalCfg.OVConfigPath + "server.conf"
			if err = config.SaveToFile("conf/openvpn-server-config.tpl",
				c.Config, destPath); err != nil {
				logs.Error(err)
			}
		}
	} else {
		logs.Error(err)
	}
}
