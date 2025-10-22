# nats-auth-callout

```
  ./tools/nsc generate nkey --account
SAACC7CFYNTN24BNYY5RWGX4VWVOKX35NYEKHVGYTSPSOCERP4HDFYPLHE
AD45NPCRRT3JN3XNVCQVIBEY7FNKCBMWBXBD3UBVDI2GALWG4BZSG6OR

```


```
./tools/nsc generate nkey --curve
SXAFZFBIBJ352B4UFKUO4MUKKKANTQYPJDL7NZFRIABVYP7FTJQ4RF7B2A
XDRKDQFCAKDH25IEWZVPA5F7674EGB5JLTAV4W5TKBLKIPYRGLEVYXIZ
```



# 生成CA
```
1.openssl genpkey -algorithm RSA -out keys/ca.key -pkeyopt rsa_keygen_bits:2048

2.openssl req -x509 -new -nodes -key keys/ca.key -sha256 -days 3650 -out keys/ca.crt -config keys/ca.conf

3.openssl x509 -req -in keys/auth.csr -CA keys/ca.crt -CAkey keys/ca.key -CAcreateserial -out keys/auth.crt -days 365 -extensions req_ext -extfile keys/auth.conf

# 生成Auth Client证书

1.openssl genpkey -algorithm RSA -out keys/auth.key -pkeyopt rsa_keygen_bits:2048

2. openssl req -new -key keys/auth.key -out keys/auth.csr -config keys/auth.conf

3.openssl x509 -req -in keys/auth.csr -CA keys/ca.crt -CAkey keys/ca.key -CAcreateserial -out keys/auth.crt -days 365 -extensions auth -extfile keys/auth.conf

# 生成Server证书
1.openssl genpkey -algorithm RSA -out keys/server.key -pkeyopt rsa_keygen_bits:2048

# 2. 生成服务器证书请求(CSR)
openssl req -new -key keys/server.key -out keys/server.csr -config keys/server.conf -batch

# 3. 使用CA证书签名服务器CSR，生成最终的服务器证书
openssl x509 -req -in keys/server.csr -CA keys/ca.crt -CAkey keys/ca.key -CAcreateserial -out keys/server.crt -days 365 -extensions req_ext -extfile keys/server.conf