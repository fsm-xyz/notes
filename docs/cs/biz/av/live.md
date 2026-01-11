# 直播

## 流和视频格式

视频解码

[参考](https://cloud.tencent.com/developer/article/1155707)

+ 基于MSE
+ 基于HLS手机端播放
+ 基于客户端播放
+ 基于flash的

RTMP推流

优点

+ 延迟小，
+ 方案比较成熟
+ 广告插入方便

缺点

+ 性能消耗大
+ 需要额外安装flash插件
+ 安全漏洞
+ 底层代码封闭

基于html5的

优点

+ 技术成本低，无需引入额外的技术
+ 性能消耗低

缺点

+ 延时性高
+ html5容易盗链
+ 方案不成熟，各家公司都有自己的标准

总结

各家公司都用自己掌握的技术，技术迁移成本大
需要年轻有想法的人去推进

### 直播协议

+ [常见直播协议](https://github.com/gwuhaolin/blog/issues/3)

+ [B站](http://easywork.xin/2018/05/05/practice-1/)

+ RTMP: 底层基于TCP，在浏览器端依赖Flash。
+ HTTP-FLV: 基于HTTP流式IO传输FLV，依赖浏览器支持播放FLV。
+ WebSocket-FLV: 基于WebSocket传输FLV，依赖浏览器支持播放FLV。WebSocket建立在HTTP之上，建立WebSocket连接前还要先建立HTTP连接。
+ HLS: Http Live Streaming，苹果提出基于HTTP的流媒体传输协议。HTML5可以直接打开播放。
+ RTP: 基于UDP，延迟1秒，浏览器不支持。

### 流媒体格式

+ FLV Flash Video
+ m3u8 ts文件

### Flash 历史

[参考](https://www.polyv.net/news/2018/12/hy0366/)

`在PC时代，Flash曾长期处于鼎盛时期，期间经历了三次高峰，分别是1999年的动画时代、2005年的Flash Video时代与2008年的Web Game时代，这三次互联网领域的高峰全都被Flash赶上。`
