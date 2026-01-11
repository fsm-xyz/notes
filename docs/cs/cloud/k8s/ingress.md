# Ingress

## cors

### 单个域名

`nginx.ingress.kubernetes.io/cors-allow-origin`设置对应的域名
```sh
metadata:
  annotations:
    nginx.ingress.kubernetes.io/configuration-snippet: ''
    nginx.ingress.kubernetes.io/cors-allow-credentials: 'true'
    nginx.ingress.kubernetes.io/cors-allow-headers: >-
      Cache-Control,Content-Type,If-Modified-Since,Keep-Alive,Origin,User-Agent,X-Requested-With
    nginx.ingress.kubernetes.io/cors-allow-methods: 'PUT, GET, POST, DELETE, OPTIONS'
    nginx.ingress.kubernetes.io/cors-allow-origin: 'xxx.com'
    nginx.ingress.kubernetes.io/enable-cors: 'true'
```

### 所有域名

类似nginx的http_origin效果

```sh
apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  name: test-cors
  annotations:
    nginx.ingress.kubernetes.io/configuration-snippet: |
      if ($request_method = 'OPTIONS') {
        more_set_headers 'Access-Control-Allow-Credentials: true';
        more_set_headers 'Access-Control-Allow-Origin: $http_origin';
        more_set_headers 'Access-Control-Allow-Methods: GET, PUT, POST, DELETE, PATCH, OPTIONS';
        more_set_headers 'Access-Control-Allow-Headers: Cache-Control,Content-Type,DNT';
        more_set_headers 'Access-Control-Max-Age: 1728000';
        more_set_headers 'Content-Type: text/plain charset=UTF-8';
        more_set_headers 'Content-Length: 0';
        return 204;
      }
      more_set_headers 'Access-Control-Allow-Credentials: true';
      more_set_headers 'Access-Control-Allow-Origin: $http_origin';
      more_set_headers 'Access-Control-Allow-Methods: GET, PUT, POST, DELETE, PATCH, OPTIONS';
      more_set_headers 'Access-Control-Allow-Headers: Cache-Control,Content-Type,DNT';
```

### 特定域名

在上面的基础上写正则， 匹配http_origin， 然后放行

+ [方案1](https://github.com/kubernetes/ingress-nginx/issues/1171#issuecomment-391988766)
+ [方案2](https://github.com/kubernetes/ingress-nginx/issues/1171#issuecomment-532158067)

### 跨域测试代码

观察控制台的报错即可

```html
<html>
<body>Hello World！</body>
</html>
<script src="https://cdn.jsdelivr.net/npm/axios/dist/axios.min.js"></script>
<script>
axios.post('http://test1-api.520yidui.com/v3/msg_lock/question', {
    identifier: this.mobile,
    certificate: this.certificate
    }, {
    withCredentials: true
    }).then((res) => {
    console.log(res)
    if (res.data.code === CodeOK) {
        localStorage.setItem('logged_in', true)
        this.$store.state.logged_in = true
        this.$router.push('guess/new')
        return
    } else {
        alert(res.data.msg)
    }
    })
</script>
```