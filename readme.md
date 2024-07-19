openssl genrsa -out server.key 204
openssl ecparam -genkey -name secp384r1 -out server.key
openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650 --addext "subjectAltName = DNS:localhost"

certificate is self-signed so may need to be trusted (by keystore on MacOS, by importing to cert store on Windows)