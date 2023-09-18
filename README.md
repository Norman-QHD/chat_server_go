## Go语言学习DEMO
### 简易聊天系统服务器

##### 基础服务端的构建.
1. 新建两个go文件 main, server. 其中main作为应用程序的主入口,用于初始化(调用)Server.server是服务端本体
2. 编写server,包含4个主要部分 a,socket监听 b,socket关闭 并且在大循环内包含 c,连接接收 d,连接处理
3. 编写Server类(使用struct)
4. 编写Server的构造函数(使用function+类型指针参数)
5. 编写Handler方法
6. 在main函数中初始化server并启动
##### User管理,管道,广播,回显
1. 创建user.go,表示一个在线用户的类型
2. 拓展两个方法: a创建用户对象 b监听对应的channel的消息
3. server类更改,新增OnlineMap和Message(广播消息)
4. 修改Handler函数,创建user和广播用户上线的逻辑.阻塞handler防止掉线
5. 新增广播消息方法,写入到server.Message
6. 新增监听广播消息的channel方法.for一直监听server.Message,找到所有的用户,然后发给所有的用户.
7. 启动Message的goroutine