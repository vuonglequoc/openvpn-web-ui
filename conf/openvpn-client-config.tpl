client

dev tun

proto {{ .Proto }}

remote {{ .ServerAddress }} {{ .Port }}

resolv-retry infinite

nobind

user nobody
group nobody

persist-key
persist-tun

ca {{ .Ca }}
cert {{ .Cert }}
key {{ .Key }}

remote-cert-tls server

#tls-auth {{ .TaKey }} 1

cipher {{ .Cipher }}
auth {{ .Auth }}

comp-lzo

verb 3

#tls-client
#lport 0
