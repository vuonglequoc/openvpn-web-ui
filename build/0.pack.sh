#!/bin/bash

set -e

PKGFILE=openvpn-web-ui.tar.gz
if [ -f $PKGFILE ]; then
    rm ../$PKGFILE
fi

time docker run \
    -v "$PWD/../":/go/src/github.com/vuonglequoc/openvpn-web-ui \
    --rm \
    -w /usr/src/myapp \
    vuonglequoc/beego:2.2.1 \
    sh -c "cd /go/src/github.com/vuonglequoc/openvpn-web-ui/ && bee version && GOFLAGS=-buildvcs=false bee pack -exr='^vendor|^data.db|^build|^README.md|^docs|^.git'"
