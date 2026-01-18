# http

大致思路

```go
http.ListenAndServe(address, handler) {
    Listen()
    for {
        Accept {
            go c.serve() {
                c.Handler.ServeHTTP()
            }
        }
    }
}
```

## Custom

```go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

```go
type ServeMux struct {
    mu    sync.RWMutex
    m     map[string]muxEntry
    hosts bool // whether any patterns contain hostnames
}

type muxEntry struct {
    h       Handler
    pattern string
}
```

```go
http.Handle()
http.HandleFunc()
```

```go
type Server struct{}
```

## FAQ

http的post上传数据问题

+ application/x-www-form-urlencoded
+ multipart/form-data
+ application/json

Form包含query parameters, body

PostForm包含body

MultipartForm包含body

ParseForm()解析query和body到Form和PostForm

MultipartForm()自动调用ParseForm(), 解析multipart/form-data下的数据到MultipartForm, 同时添加到Form和PostForm
