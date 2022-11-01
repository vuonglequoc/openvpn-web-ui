#!/bin/sh

set -e

if [ ! -f $OPENVPN/.provisioned ]; then
  echo "Preparing certificates"
  /opt/scripts/generate_ca_and_server_certs.sh
  touch $OPENVPN/.provisioned
fi

cd /opt/openvpn-gui

mkdir -p db

./openvpn-web-ui

echo "Starting!"
