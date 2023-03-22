echo '> 生成node.key'
openssl genrsa -des3 -out node.key 2048
echo '> 生成不需要密码的node.key'
openssl rsa -in node.key -out node.key
echo '> 生成ca的crt'
openssl req -new -x509 -key node.key -out ca.crt -days 365
echo '> 生成csr'
openssl req -new -key node.key -out node.csr
echo '> 生成crt'
openssl x509 -req -days 365 -in node.csr -CA ca.crt -CAkey node.key -CAcreateserial -out node.crt
echo '> 证书合并'
cat node.key node.crt > node.pem