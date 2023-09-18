## Go语言学习DEMO
### 简易聊天系统服务器

##### 基础服务端的构建.
1. 新建两个go文件 main, server. 其中main作为应用程序的主入口,用于初始化(调用)Server.server是服务端本体
2. 编写server,包含4个主要部分 a,socket监听 b,socket关闭 并且在大循环内包含 c,连接接收 d,连接处理
3. 编写Server类(使用struct)
4. 编写Server的构造函数(使用function+类型指针参数)
5. 编写Handler方法
6. 在main函数中初始化server并启动