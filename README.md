# OpenVPN-web-ui

[![Build Status](https://jenkins.kinguda.com/buildStatus/icon?job=openvpn-web-ui)](https://jenkins.kinguda.com/buildStatus/icon?job=openvpn-web-ui)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fvuonglequoc%2Fopenvpn-web-ui.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fvuonglequoc%2Fopenvpn-web-ui?ref=badge_shield)

## Summary

OpenVPN server web administration interface.

Goal: create quick to deploy and easy to use solution that makes work with small OpenVPN environments a breeze.

If you have docker and docker-compose installed, you can jump directly to [installation](#Prod).

![Status page](docs/images/preview_status.png?raw=true)

Please note this project is in alpha stage. It still needs some work to make it secure and feature complete.

## Motivation

## Features

* status page that shows server statistics and list of connected clients
* easy creation of client certificates
* ability to download client certificates as a zip package with client configuration inside
* log preview
* modification of OpenVPN configuration file through web interface

## Screenshots

[Screenshots](docs/screenshots.md)

## Usage

After startup web service is visible on port 8080. To login use the following default credentials:

* username: admin
* password: b3secure (this will be soon replaced with random password)

Please change password to your own immediately!

### Prod

Requirements:
* docker and docker-compose
* on firewall open ports: 1194/udp and 8080/tcp

Execute commands

    curl -O https://raw.githubusercontent.com/vuonglequoc/openvpn-web-ui/master/docs/docker-compose.yml
    docker-compose up -d

It starts two docker containers. One with OpenVPN server and second with OpenVPNAdmin web application. Through a docker volume it creates following directory structure:

    .
    ├── docker-compose.yml
    └── openvpn-data
        ├── openvpn
        │   ├── client-configs
        │   │   ├── files
        │   │   │   └── client_*.ovpn
        │   │   └── keys
        │   │       └── client_*.crt
        │   │── pki
        │   │   ├── private
        │   │   │   ├── client_*.key
        │   │   │   └── server.key
        │   │   ├── reqs
        │   │   │   ├── client_*.req
        │   │   │   └── server.req
        │   │   ├── dh2048.pem
        │   │   ├── ca.crt
        │   │   ├── server.crt
        │   │   ├── ta.key
        │   │   ├── openssl-easyrsa.cnf
        │   │   └── safessl-easyrsa.cnf
        │   ├── log
        │   │   ├── ipp.txt
        │   │   ├── openvpn.log
        │   │   └── openvpn-status.log
        │   ├── server.conf
        │   └── vars
        ├── ca_server
        │   ├── pki
        │   │   ├── certs_by_serial
        │   │   │   └── *.pem
        │   │   ├── issued
        │   │   │   ├── client_*.crt
        │   │   │   └── server.crt
        │   │   ├── private
        │   │   │   └── ca.key
        │   │   ├── reqs
        │   │   │   ├── client_*.req
        │   │   │   └── server.req
        │   │   ├── ca.crt
        │   │   ├── index.txt
        │   │   ├── index.txt.attr
        │   │   ├── index.txt.attr.old
        │   │   ├── index.txt.old
        │   │   ├── index_ok.txt
        │   │   ├── serial
        │   │   └── serial.old
        │   └── vars
        └── db
            └── data.db

### Dev

Requirements:
* [golang 1.20.5](https://hub.docker.com/_/golang)
* [beego v2.1.0](https://github.com/beego/beego)
* [bee v2.0.5](https://github.com/beego/bee)

Execute commands:

    go get github.com/vuonglequoc/openvpn-web-ui
    cd $GOPATH/src/github.com/vuonglequoc/openvpn-web-ui
    bee run -gendoc=true

### Source code structure

    .
    ├── build
    │   ├── Dockerfile
    │   └── [build scripts]
    ├── conf
    │   ├── app.conf
    │   ├── openvpn-client-config.tpl
    │   └── openvpn-server-config.tpl
    ├── controllers                 # MVC
    ├── docs
    │   ├── docker-compose.yml
    │   └── [documents]
    ├── lib                         # Lib for controllers
    ├── models                      # MVC
    ├── routers                     # Application routes
    ├── static                      # CSS, Img, JS
    ├── swagger                     # RESTful APIs (beego generated)
    ├── vendor
    ├── view                        # MVC - AdminLTE
    │   ├── common
    │   │   ├── alert.html
    │   │   ├── footer.html
    │   │   ├── fvalid.html
    │   │   └── [header].html
    │   ├── layout
    │   │   └── base.html
    │   └── [page].html
    ├── main.go
    ├── go.mod
    ├── go.sum
    ├── LICENSE
    └── README.md

### Compiled structure

    .
    ├── conf
    │   ├── app.conf
    │   ├── openvpn-client-config.tpl
    │   └── openvpn-server-config.tpl
    ├── db
    │   └── data.db
    ├── static
    ├── swagger
    ├── view                        # MVC
    │   ├── common
    │   │   ├── alert.html
    │   │   ├── footer.html
    │   │   ├── fvalid.html
    │   │   └── [header].html
    │   ├── layout
    │   │   └── base.html
    │   └── [page].html
    ├── openvpn-web-ui              # main app
    └── LICENSE

## Important Note

### Management interface

OpenVPNAdmin will manage OpenVPN daemon via management api.

https://openvpn.net/community-resources/how-to/

In order to enable management api for OpenVPN daemon, we need to add below config to the config file `server.conf`:

`management 0.0.0.0 2080`

In the `Settings` of OpenVPNAdmin, update `Management interface address` with IP of OpenVPN daemon and same port as above (2080).

### Logging

OpenVPNAdmin will read OpenVPN daemon log from `/etc/openvpn/log/openvpn.log`.

In order to enable this log for OpenVPN daemon, we need to add below config to the config file `server.conf`:

`log-append /etc/openvpn/log/openvpn.log`

### SSL

Added SSL Support by adding HTTPS config in `app.conf`

```
appname = openvpn-web-ui
httpport = 8080
runmode = dev
EnableGzip = true
EnableAdmin = true
sessionon = true
CopyRequestBody = true

HTTPSCertFile = /opt/certs/vpn.example.com/cert.pem
HTTPSKeyFile = /opt/certs/vpn.example.com/privkey.pem
HTTPSPort = 443
EnableHTTPS = true

DbPath = "./data.db"
```

## Todo

* add unit tests
* add option to modify certificate properties
* generate random admin password at initialization phase
* add versioning
* add automatic ssl/tls (check how [ponzu](https://github.com/ponzu-cms/ponzu) did it)

## Remarks

### Vendoring
To update dependencies from GOPATH:

`go mod vendor`

### Template

- [AdminLTE 2.3.7](https://github.com/ColorlibHQ/AdminLTE) - dashboard & control panel theme. Built on top of Bootstrap.
- Bootstrap 3.3.7
- FontAwesome 5.15.3
- Ionicons 2.0.0
- iCheck 1.0.2

## License

This project uses [MIT license](LICENSE)

[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fvuonglequoc%2Fopenvpn-web-ui.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fvuonglequoc%2Fopenvpn-web-ui?ref=badge_large)
