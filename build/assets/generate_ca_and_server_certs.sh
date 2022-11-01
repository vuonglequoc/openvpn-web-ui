#!/bin/sh -e

EASY_RSA=/usr/share/easy-rsa

echo "Configuration ENV"

echo $SERVER_NAME
echo $SERVER_EMAIL
echo $SERVER_COUNTRY
echo $SERVER_PROVINCE
echo $SERVER_CITY
echo $SERVER_ORG
echo $SERVER_OU
echo $SERVER_CN


echo "Generating CA cert"

# Preparing a Public Key Infrastructure Directory
cd $CA_SERVER
$EASY_RSA/easyrsa init-pki

# Prepare config params
cp /opt/scripts/vars.template $CA_SERVER/vars

# Creating a Certificate Authority
$EASY_RSA/easyrsa --batch --req-cn=$SERVER_NAME build-ca nopass


echo "Generating Server cert"

# Creating a PKI for OpenVPN
cd $OPENVPN
echo "set_var EASYRSA_ALGO ec" > $OPENVPN/vars
echo "set_var EASYRSA_DIGEST sha512 " >> $OPENVPN/vars

$EASY_RSA/easyrsa init-pki

# Creating an OpenVPN Server Certificate Request and Private Key
# echo -e "$SERVER_NAME" | $EASY_RSA/easyrsa gen-req $SERVER_NAME nopass
openssl genrsa -out $OPENVPN/pki/private/$SERVER_NAME.key 2048
openssl req -new -key $OPENVPN/pki/private/$SERVER_NAME.key -out $OPENVPN/pki/reqs/$SERVER_NAME.req \
    -subj /emailAddress="$SERVER_EMAIL"/C="$SERVER_COUNTRY"/ST="$SERVER_PROVINCE"/L="$SERVER_CITY"/O="$SERVER_ORG"/OU="$SERVER_OU"/CN="$SERVER_CN"
# key: $OPENVPN/pki/private/SERVER_NAME.key
# req: $OPENVPN/pki/reqs/SERVER_NAME.req

# Signing the OpenVPN Server’s Certificate Request
cd $CA_SERVER
$EASY_RSA/easyrsa import-req $OPENVPN/pki/reqs/$SERVER_NAME.req $SERVER_NAME
echo -e "yes" | $EASY_RSA/easyrsa sign-req server $SERVER_NAME
# cert: $CA_SERVER/pki/issued/SERVER_NAME.crt

# Copy Signed Certificate
cp $CA_SERVER/pki/issued/$SERVER_NAME.crt $OPENVPN/pki/$SERVER_NAME.crt
cp $CA_SERVER/pki/ca.crt $OPENVPN/pki/ca.crt
# cert: $OPENVPN/pki/SERVER_NAME.crt
# cert: $OPENVPN/pki/ca.crt
