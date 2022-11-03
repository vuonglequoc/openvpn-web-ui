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

#ca {{ .Ca }}
#cert {{ .Cert }}
#key {{ .Key }}

remote-cert-tls server

#tls-auth {{ .TaKey }} 1

key-direction 1

cipher {{ .Cipher }}
auth {{ .Auth }}

#comp-lzo

verb 3

{{ .ExtraClientOptions }}

; script-security 2
; up /etc/openvpn/update-systemd-resolved
; down /etc/openvpn/update-systemd-resolved
; down-pre
; dhcp-option DOMAIN-ROUTE .
