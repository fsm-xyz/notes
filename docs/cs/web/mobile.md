# mobile

## 移动端安卓和iphone的click时间区别

>- 问题描述：
  在代码中使用.on()给元素绑定click事件，安卓和模拟器中表现正常，iphone表现异常
>- 原因分析：
  由于在iphone中Safri认为一个元素有这个属性才是一个可点击区域，才可以具有click事件。
>- 解决方案：
  1.给对应的节点加上(cursor:pointer)
  2.给对应的节点加上(onclick="")

## APP类型

  Web App             套一个壳子，里面通过web技术实现
  Hybrid App          介于2者
  Native APP          使用原生的技术开发
  
### Hybird App

  appMobi, PhoneGap, WeX5, AppCan, Rexsee, Appcelerator, Kerkee, APICloud
       NativeScript, Kinvey, ExMobi

### 跨平台

  Flutter
  React Native
  Weex
  QT

### 热更新

  JSPath

## APP和网页

### Android打开网页

- 系统自带的浏览器访问
- WebView控件

### IOS打开网页

- 调用Safri(跳转)
- 调用Safri(不跳转)
- WKWebView

### 网页打开APP

调用特定的url
