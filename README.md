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

##### 群聊基本功能
1. 修改Server.go 完善 Handler 针对当前客户端的消息读取业务,开一个go程
2. 创建缓冲区并开始循环读取
3. 如果读取到0 客户端合法关闭
4. 如果是合法消息,把消息拿到,去掉回车再进行一次广播.

##### 做一些封装
1. 在user中新增server的指针关联.
2. new User 增加一个server指针的赋值
3. 增加Online和Offline方法
4. 由User处理Message 使用DoMessage
5. 整合Server和User,冗余的替换掉.
6. 测试上线,下线,消息的业务(回调函数测试)

##### 查询在线用户列表
1. 修改user.go添加一个SendMessage消息(单发)
2. DoMessage 加上"who"指令的处理
3. 拼装who返回结果的string,回发给当前user

##### 修改用户名
1. 在DoMessage中添加else if的处理,满足rename|开头的内容,则视为修改用户名请求
2. 通过字符串split截取出来要修改的用户名
3. 判断用户名是否存在,如果存在则直接返回消息说已存在,不进行后续处理.
4. 从server的map里面通过key移除,添加新的key和user(就是改了引用的名字)
5. 修改自身对象(user)的Name字段
6. 回显消息给自己

##### 超时强制踢出功能
1. 在每个用户的handler()的goroutine中,添加用户是否活跃的channel,一旦有消息,就像channel发送消息
2. 只要用户可度任何的消息,处理完了消息以后,让isLive写一个true
3. 监听isAlive和定时器的监听. isLive写在上面,只要活跃就不会执行timer,到timer的时候就强制踢出

##### 私聊功能
1. 判断消息格式是否满足
2. 判断用户名是否存在
3. 如果存在会找到对方用户名对象
4. 获取要发送的消息的内容
5. 通过对方的对象,发送给对方的用户