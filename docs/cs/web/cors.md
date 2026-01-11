# 跨域

+ ajax跨域
+ Cookie跨域
+ iframe跨域
+ LocalStorage跨域

## 背景

同源策略

+ 协议相同
+ 域名相同
+ 端口相同

同源政策的目的，是为了保证用户信息的安全，防止恶意的网站窃取数据。

## 跨域解决方案

+ 代理
+ 图片ping或script标签(src属性)
+ JSONP
+ Websocket
+ CORS
+ window.name+iframe
+ window.postMessage()
+ document.domain跨子域

## CORS请求分类

+ 简单请求
不做预检
+ 复杂请求
做预检，浏览器会先发送OPTIONS请求,通过则继续

### CORS解决方案

+ nginx方案
+ 后端服务器直接处理

主要设置如下4个个字段

1. Access-Control-Allow-Origin
2. Access-Control-Allow-Methods
3. Access-Control-Allow-Headers
4. Access-Control-Allow-Credentials

### Nginx方案

```bash
location / {
add_header Access-Control-Allow-Origin $http_origin;
add_header Access-Control-Allow-Headers X-Requested-With,Content-Type;
add_header Access-Control-Allow-Methods GET,POST,OPTIONS;
add_header Access-Control-Allow-Credentials true;

# 复杂请求的预检
if ($request_method = 'OPTIONS') {
return 204;
}
}
```

### 浏览器

浏览器会自动加这几个字段对应的东西
Access-Control-Request-Headers: access-control-request-origin,content-type
Access-Control-Request-Method: POST
Origin: <http://xxx.com>

### 验证cors

```sh
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

### 常见错误

1.Response to preflight request doesn't pass access control check: The value of the 'Access-Control-Allow-Origin' header in the response must not be the wildcard '*' when the request's credentials mode is 'include'.

浏览器发送OPTIOSN操做没有验证通过, 使用了Cookie传输设置, ORIGIN必须指定对应的域,不能用*
If you want to allow credentials then your Access-Control-Allow-Origin must not use *. You will have to specify the exact protocol + domain + port.

2.Request header field access-control-allow-origin is not allowed by access-control-allow-headers in preflight response.

浏览器传输的header和服务器的Access-Control-Allow-Headers 不一致

3.Response to preflight request doesn't pass access control check: No 'Access-Control-Allow-Origin' header is present on the requested resource

如果http返回非2xx，则丢掉返回Header中的 Access-Control-Allow-Origin, 报跨域问题, 解决这个问题设置http默认返回200，业务中使用自定义使用json数据返回码

 4.refused 拒绝设置默写header等
browsers do not allow head has set some security risks, such as cookie, host, referer, etc. So, do not use the browser to parse the line head. It can be used on the server side set of agency sent.

## 参考

[Mozila](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Access_control_CORS)
[阮一峰](http://www.ruanyifeng.com/blog/2016/04/cors.html)
[CSDN](https://blog.csdn.net/ligang2585116/article/details/73072868)
