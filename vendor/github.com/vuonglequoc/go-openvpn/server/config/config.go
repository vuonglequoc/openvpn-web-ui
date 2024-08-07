package config

import (
	"bytes"
	"html/template"
	"io/ioutil"
)

var defaultConfig = Config{
	Port:                1194,
	Proto:               "udp",
	Ca:                  "ca.crt",
	Cert:                "server.crt",
	Key:                 "server.key",
	Dh:                  "dh2048.pem",
	Server:              "10.8.0.0 255.255.255.0",
	IfconfigPoolPersist: "ipp.txt",
	DNSServerOne:        "8.8.8.8",
	DNSServerTwo:        "8.8.4.4",
	Keepalive:           "10 120",
	TaKey:               "ta.key",
	Cipher:              "AES-256-CBC",
	Auth:                "SHA256",
	MaxClients:          100,
	Management:          "0.0.0.0 2080",
	ExtraServerOptions:  "",
	ExtraClientOptions:  "",
}

//Config model
type Config struct {
	Port int
	Proto string

	Ca string
	Cert string
	Key string

	Dh string

	Server string

	IfconfigPoolPersist string

	DNSServerOne string
	DNSServerTwo string

	Keepalive string

	TaKey string

	Cipher string
	Auth string

	MaxClients int

	Management string

	ExtraServerOptions string
	ExtraClientOptions string
}

//New returns config object with default values
func New() Config {
	return defaultConfig
}

//GetText injects config values into template
func GetText(tpl string, c Config) (string, error) {
	t := template.New("config")
	t, err := t.Parse(tpl)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	t.Execute(buf, c)
	return buf.String(), nil
}

//SaveToFile reads teamplate and writes result to destination file
func SaveToFile(tplPath string, c Config, destPath string) error {
	template, err := ioutil.ReadFile(tplPath)
	if err != nil {
		return err
	}

	str, err := GetText(string(template), c)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(destPath, []byte(str), 0644)
}
