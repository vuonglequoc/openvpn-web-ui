port {{ .Port }}
proto {{ .Proto }}

dev tun

ca {{ .Ca }}
cert {{ .Cert }}
key {{ .Key }}

dh {{ .Dh }}

server {{ .Server }}

ifconfig-pool-persist {{ .IfconfigPoolPersist }}

push "route {{ .Server }}"

push "redirect-gateway def1 bypass-dhcp"

push "dhcp-option DNS {{ .DNSServerOne }}"
push "dhcp-option DNS {{ .DNSServerTwo }}"

keepalive {{ .Keepalive }}

tls-auth {{ .TaKey }} 0
#tls-crypt {{ .TaKey }}

key-direction 0

cipher {{ .Cipher }}
auth {{ .Auth }}

askpass /etc/openvpn/pki/passphrase.txt

#comp-lzo

max-clients {{ .MaxClients }}

user nobody
group nobody

persist-key
persist-tun

status /etc/openvpn/log/openvpn-status.log

log-append /etc/openvpn/log/openvpn.log

verb 3

mute 10

management {{ .Management }}

{{ .ExtraServerOptions }}
