#!/bin/bash -e

EASY_RSA=/usr/share/easy-rsa

echo "Generating CA cert"

# Preparing a Public Key Infrastructure Directory
cd $CA_SERVER
$EASY_RSA/easyrsa init-pki

# Prepare config params
cp $CA_SERVER/scripts/vars.template $CA_SERVER/vars

# Creating a Certificate Authority
$EASY_RSA/easyrsa --batch build-ca nopass


echo "Generating Server cert"

SERVER_NAME=Server

# Creating a PKI for OpenVPN
cd $OPENVPN
echo "set_var EASYRSA_ALGO ec" > $OPENVPN/vars
echo "set_var EASYRSA_DIGEST sha512 " >> $OPENVPN/vars

$EASY_RSA/easyrsa init-pki

# Creating an OpenVPN Server Certificate Request and Private Key
echo -e "server" | $EASY_RSA/easyrsa gen-req $SERVER_NAME nopass
# key: $OPENVPN/pki/private/SERVER_NAME.key
# req: $OPENVPN/pki/reqs/SERVER_NAME.req

# Signing the OpenVPN Server’s Certificate Request
cd $CA_SERVER
$EASY_RSA/easyrsa import-req $OPENVPN/pki/reqs/$SERVER_NAME.req $SERVER_NAME
echo -e "yes" | $EASY_RSA/easyrsa sign-req server $SERVER_NAME
# cert: $CA_SERVER/pki/issued/SERVER_NAME.crt

cp $CA_SERVER/pki/issued/$SERVER_NAME.crt $OPENVPN/pki/$SERVER_NAME.crt
cp $CA_SERVER/pki/ca.crt $OPENVPN/pki/ca.crt
# cert: $OPENVPN/pki/SERVER_NAME.crt
# cert: $OPENVPN/pki/ca.crt
