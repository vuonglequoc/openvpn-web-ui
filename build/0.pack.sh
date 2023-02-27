#!/bin/bash

set -e

PKGFILE=openvpn-web-ui.tar.gz
rm ../$PKGFILE

time docker run \
    -v "$PWD/../":/go/src/github.com/vuonglequoc/openvpn-web-ui \
    --rm \
    -w /usr/src/myapp \
    vuonglequoc/beego:2.0.7 \
    sh -c "cd /go/src/github.com/vuonglequoc/openvpn-web-ui/ && bee version && bee pack -exr='^vendor|^data.db|^build|^README.md|^docs|^.git'"
