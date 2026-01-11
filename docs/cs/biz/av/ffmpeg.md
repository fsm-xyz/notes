# ffmpeg

## 参数

```sh
-an  # 不推送音频
-re  # 循环推送
-rtsp_transport # udp:默认 数据包丢失或卡顿， tcp: 更大的延迟，http, udp_multicast:  使用多播 UDP
```

# 编解码

## 

```sh
# 查看所有支持的编解码器（包含编码器和解码器）
ffmpeg -codecs

# 查看所有支持的格式
ffmpeg -formats

# 查看所有编码器
ffmpeg -encoders

# 查看特定类型的编码器（如视频编码器）
ffmpeg -encoders | grep Video

# 查看特定编码器的详细信息（以 libx264 为例）
ffmpeg -h encoder=libx264

# 查看所有解码器
ffmpeg -decoders

# 查看特定类型的解码器（如音频解码器）
ffmpeg -decoders | grep Audio

# 查看特定解码器的详细信息（以 aac 为例）
ffmpeg -h decoder=aac

-preset     ultrafast~veryslow  preset 编解码复杂度，越大越快
-cpu-used   0-9                 数字越大越快
-threads    线程数

-crf 动态调整bitrate    越大越不清晰
-b:v 直接指定bitrate
```