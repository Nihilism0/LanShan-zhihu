# 蓝山Web寒假双人考核-知乎

### 接口文档📃

[蓝山知乎接口](https://console-docs.apipost.cn/preview/648a7969b340d643/bd974949c5b66514)

------

前端部署地址 https://zhihu.madeindz.work/

后端部署地址 https://gogo.madeindz.work/

------

后端项目github地址:

## 前言😆

和前端同学一起花了不少时间完成的这个项目，虽然还是有许多不足的地方，但还是有我们自己的努力在里面。文档中有一些功能的说明，希望学长们可以看看。🥰

## **技术栈**💫

**Gin框架**

**JWT**

**Cors**

**Websocket**

**NSQ**

**SMS**

**Viper**

**令牌桶/juju包**

**Zap**

**MD5**

**Mysql**

**Redis**

## 已实现功能列表⛳

*粗体为基础功能要求*

- **用户注册**
- **用户登录**
- **发布问题**
- **回答问题**
- **发布文章**
- **评论回答和评论文章**
- **我的（个人信息栏）**
- **用户信息更改（昵称，密码, 性别 ,签名 等）**
- **我的收藏(收藏夹)**
- 热榜
- 用户密码加盐加密
- 用户登录有短信登录、邮箱登录、手机登录
- 验证码（登录，注册，修改密码）
- 用户状态保存使用 JWT
- 搜索功能（搜索问题或文章）
- 点赞功能（赞同问题、文章、评论）
- 关注功能（关注的人提问，回答，点赞会得到通知）
- 盐选会员（只有氪金才能看完整的知乎文章）
- 用户浏览记录
- 将项目部署上线（包括前端和后端的项目，也就是登录你的网站能够像正常的网站一样访问）
- 使用 https 加密
- 私信聊天
- 商城(盐选会员获取渠道)
- 限流(令牌桶)

## 亮点✨

### 短信业务💌																																																																													

外接**腾讯云SMS短信业务**,通过随机生成六位数的函数获得验证码信息, 发送至用户的手机并Set进Redis数据库, 并设置TTL为3分钟, 在需要验证时Get并校对.

此短信业务在登录系统, 注册系统, 修改密码时会用到.

![image-20230111163221151](C:\Users\Ambrose\AppData\Roaming\Typora\typora-user-images\image-20230111163221151.png)

### NSQ+Websocket📨

使用NSQ进行做消息队列,消息队列把业务流程中的非关键流程异步化，从而显著降低业务请求的响应时间

使用带状态的协议websocket让服务器主动推送消息

![image-20230109205836127](C:\Users\Ambrose\AppData\Roaming\Typora\typora-user-images\image-20230109205836127.png)

生产者(关注的人有动作便进行)

![image-20230111162823603](C:\Users\Ambrose\AppData\Roaming\Typora\typora-user-images\image-20230111162823603.png)

消费者(main函数中起线程,监听nsq管道)

![image-20230111162847525](C:\Users\Ambrose\AppData\Roaming\Typora\typora-user-images\image-20230111162847525.png)

通过用户是否连接websocket判断read是0(未读)还是1(已读)

![image-20230110215548761](C:\Users\Ambrose\AppData\Roaming\Typora\typora-user-images\image-20230110215548761.png)

若在线会进行即时推送,并把消息read设置为1

![image-20230110224735543](C:\Users\Ambrose\AppData\Roaming\Typora\typora-user-images\image-20230110224735543.png)

用户若点进关注铃铛按钮,看见未读信息并自动将其设置为1

### HOT值比重🔥

热榜例如回答的**HOT**值由点赞数和评论数通过一定比重相加排列而成,不仅仅只由单一指标决定

![image-20230105154510312](C:\Users\Ambrose\AppData\Roaming\Typora\typora-user-images\image-20230105154510312.png)

### 收藏夹及其隐私🗃️

用户收藏可以设置**收藏夹**, 使用外键约束的方式关联 收藏实例表.

每个收藏夹都设计了Private隐私选项,1为隐私,0为公开(默认).

用户随时可以更改收藏夹名和其隐私性和描述,也会进行逻辑上的收藏夹重名检查

![image-20230105154901106](C:\Users\Ambrose\AppData\Roaming\Typora\typora-user-images\image-20230105154901106.png)

### WebSocket聊天👩🏼‍❤️‍👨🏻

设置**防骚扰模式**: 若单方面发送消息三条以上便限制其发送,

历史条数通过Redis储存,消息实例通过MongoDB储存

![image-20230106182612192](C:\Users\Ambrose\AppData\Roaming\Typora\typora-user-images\image-20230106182612192.png)

判断对方是否在线,原理是在总的Clients管理结构体中查找ID是否有SendID,以此来决定信息

![image-20230105160908380](C:\Users\Ambrose\AppData\Roaming\Typora\typora-user-images\image-20230105160908380.png)

### **商城**🏪

用户拥有**GO币**🪙(获得GO币的渠道是发布问题+3,回答+5,文章+5,评论+1),可以在商城购买VIP等

每件商品有描述和销量

![image-20230106191028746](C:\Users\Ambrose\AppData\Roaming\Typora\typora-user-images\image-20230106191028746.png)

有shoplist表,可以查看用户的购买历史,用作用户查询和dao层逻辑判断

![image-20230107013316069](C:\Users\Ambrose\AppData\Roaming\Typora\typora-user-images\image-20230107013316069.png)

### 令牌桶限流🪣

使用令牌桶限流策略 ,限制请求数量, 取令牌集中在发布问题文章等地方,防止恶意攻击, 加强服务的稳定性

![image-20230112123237992](C:\Users\Ambrose\AppData\Roaming\Typora\typora-user-images\image-20230112123237992.png)

![image-20230112121143012](C:\Users\Ambrose\AppData\Roaming\Typora\typora-user-images\image-20230112121143012.png)

## 难点与收获😿

这些难点有些是因为它就是很难,有些是因为知识层面不够才难,知道了后便不算难点,我遇到的大多为后者

### NSQ+Websocket

使用websocket这个初学陌生的协议来监听是否发送信息等,运用管道和goroutine

第一次接触,感觉比较难

### http / https

使用https加密需要的SSL证书必须涵盖此网站,而且端口为80或者**443**才是安全端口,内部自己设置的端口会判断为不安全

### CORS跨域问题

前端与我对接时发生了跨域问题,折腾很久发现是前端发送请求会有两次,第一次是OPTIONS请求定位服务器,第二次是真正的请求,所以CORS就要有一个逻辑,当Request是OPTIONS方法就要给其200,才不会报错

![image-20230107142119355](C:\Users\Ambrose\AppData\Roaming\Typora\typora-user-images\image-20230107142119355.png)

### MySql设置问题

使用gorm进行外键约束后,其默认删除和改变主键的模式都是RESTRICT,这不好,当要删除主键所在表的数据,会被外键约束报错,而可以把这个模式改为CASCADE,当主键被改时其外键也一样被改,主键删除后外键也删除. 用户进行改名操作时十分有用,不需要所有的表都去update username

![image-20230107142924873](C:\Users\Ambrose\AppData\Roaming\Typora\typora-user-images\image-20230107142924873.png)



