# ait

此项目为个人的业余项目，主要目的是想以全栈的方式从前端到后端，从移动端到桌面，提供一个具有社交功能的助手服务，比如：扫码翻译、名片识别、电子名片交换、即时通讯、语音识别、二维码扫描等功能。当前项目为服务端，主要功能是提供GraphQL和WebSocket的数据接口。

项目地址如下： [https://github.com/JermineHu/ait](https://github.com/JermineHu/ait) 此项目为个人的业余项目主要包含如下功能：

## 服务端介绍：

1. 个人认为采用REST API设计接口的在某些调用上很不优雅，所以比较推崇graphql作为接口描述，该项目主要提供graphql服务，并直接支持http2.0；
2. 内部web框架基于echo，能保证轻量级和最小化编译；
3. 提供websocket服务，用于服务端推送数据到客户端，主要是websocket跨平台，对web前端友好；
4. 后期将增加webrtc服务，主要有两方面目的，一个是实时视频语音通话（一对一），另一个就是视频会议或直播（一对多，借助cdn或者nginx通过rtmp或HLS协议在客户端获取视频资源，已经做了了解并确认了可行性）；
5. 数据库存储将选用TIDB，因为看他们的roadmap，将同时支持pg和mysql协议，并已经支持json，后续还会加入文档存储很期待，毕竟tidb也是用go和rust写的，都是我比较喜欢的；
6. go的依赖包采用glide，不解释它的强大；
7. 微服务之间采用grpc调用，序列化采用protobuff，因为它可以跨语言调用，设计完美；
8. 网站静态化将采用强大的hugo；
9. 为了提高网站的相应和处理速度，支持了消息队列Kafka；
10. 搜索引擎用ES；
11. 服务发现用consul;

## 移动端产品：

1. 采用跨平台方案 React Native程序编写
2. 序列化采用protobuff
3. 查询采用grahql实现
4. 客户端通讯用HTTP2.0，将需要的数据也一并推送给客户端
5. 原生程序Android采用Kotlin实现，IOS采用Swift
6. 采用websocket进行广播通知

## 桌面程序：

1. 采用Electron编写跨平台的桌面应用
2. 序列化采用protobuff
3. 查询采用grahql实现
4. 客户端通讯用HTTP2.0，将需要的数据也一并推送给客户端

   5.采用websocket进行广播通知

## DevOpts：

1. CI使用Drone
2. 使用vagrant统一开发环境
3. 自动编排和调度使用rancher
4. 支持docker的基础镜像使用coreos或者rancheros
5. 环境统一采用yml来配置，只需要执行以下docker-compose即可启动一个完整的开发环境，比如redis、mongodb等服务，直接一步到位；
6. 反向代理 使用traefik；
7. git服务器采用gitea；

