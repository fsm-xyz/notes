## curl

-k          忽略证书
-i
-v
--cacert    使用证书

curl --cacert my-ca.crt https://xxx.com

### https

展开证书内容

openssl x509 -in my-ca.crt -text