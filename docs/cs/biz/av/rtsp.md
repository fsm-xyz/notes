# RTSP

主要用于视频监控

rtsp://摄像头用户名:密码@地址：端口 服务器上地址参数

## 摄像头流

```sh
# 主
3 rtsp://admin:xxx@xxx:554/h264/ch1/main/av_stream
# 旁路
4 rtsp://admin:xxx@xxx:554/Streaming/Channels/101?transportmode=unicast
```

## 转换

由于现在主流浏览器的播放器不能直接播放rtsp流，需要通过特殊手段播放

+ rtsp转别的
+ rtsp插件