# JSON

## 解析

当进行map[string]interface{}解析的时候，json字符串里面的数字默认当做json.Number类型
可以自定义decode, 并且d.UseNumber()进行处理，避免float64溢出

## tag

```go
type S struct {
    A int `json:"a,string"`
    B int `json:"b,-"`
    C int `json:"c,omitempty"`
}
```

- 不需要这个字段
omitempty 为空, 忽略
可以定义tag为string, 转换成int

## 自定义

定义自己的实现

- MarshalJSON
- UnmarshalJSON

### JSON处理

对一个空map(没有具体指向)进行marshal，得到null
对一个对一个有指向，没内容的map得到{}
空指针或者引用，go里面就会转成null

```go
func main() {
         var a, b map[string]interface{}
         b = make(map[string]interface{})
         aa, _ := json.Marshal(a)
         bb, _ := json.Marshal(a)
         fmt.Println(string(aa), string(bb))
}
// result : null {}
```

[参考](https://colobu.com/2017/06/21/json-tricks-in-Go)
