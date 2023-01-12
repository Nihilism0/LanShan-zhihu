package api

import (
	"CSAwork/service"
	"CSAwork/utils/middleware"
	"github.com/gin-gonic/gin"
	"time"
)

func InitRouter() {
	r := gin.Default()
	r.Use(middleware.CORS())
	r.Use(middleware.RateLimitMiddleware(1*time.Millisecond, 20))
	r.POST("/api/getcode", GetCode)                       //手机获得验证码
	r.POST("/api/register", Register)                     // 注册(需要手机验证码)
	r.POST("/api/accountlogin", AccountLogin)             //账号密码登录
	r.POST("/api/emaillogin", EmailLogin)                 //邮箱密码登录
	r.POST("/api/phonepasswordlogin", PhonePasswordLogin) //手机号密码登录
	r.POST("/api/phonelogin", PhoneLogin)                 //手机号验证码登录
	UserRouter := r.Group("/api/user")
	{
		UserRouter.Use(middleware.JWTAuthMiddleware())
		UserRouter.PUT("/informationmodify", InformationModify)   //更改用户个人信息
		UserRouter.GET("/getcode", TokenGetCode)                  //更改密码的手机获得验证码
		UserRouter.PUT("/changepassword", PasswordModify)         //更改密码
		UserRouter.GET("/getuserinformation", GetUserInformation) //获得用户信息
		UserRouter.GET("/getuserheadphoto", GetUserHeadPhoto)     //根据名字获得头像
	}
	QARouter := r.Group("/api/qa") //QA是Question和Answer
	{
		//首字母为q或a表明是问题还是回答(文章) 后面跟动词
		QARouter.Use(middleware.JWTAuthMiddleware())
		QARouter.POST("/qcreate", Qcreate)                              //发布问题
		QARouter.GET("/qsubmited", Qsubmited)                           //查看某人发表过的问题及其解答
		QARouter.POST("/acreate", Acreate)                              //发布回答
		QARouter.GET("/asubmited", Asubmited)                           //查看某人发布过的回答及其问题和评论
		QARouter.PUT("/qmodify", Qmodify)                               //修改问题
		QARouter.PUT("/amodify", Amodify)                               //修改回答
		QARouter.DELETE("/qdelete", Qdelete)                            //删除问题
		QARouter.DELETE("/adelete", Adelete)                            //删除回答
		QARouter.POST("/acomment", Acomment)                            //评论回答(楼中楼)
		QARouter.GET("/getquestionfromanswer", GetQuestionFromAnswer)   //根据回答查问题
		QARouter.GET("/seequestioninformation", SeeQuestionInformation) //看某一问题信息
		QARouter.GET("/randquestions", RandQuestions)                   //返回全部问题
	}
	ArticleRouter := r.Group("/api/article")
	{
		ArticleRouter.Use(middleware.JWTAuthMiddleware())
		ArticleRouter.POST("/createarticle", CreateArticle)                 //发布文章
		ArticleRouter.DELETE("/articledelete", ArticleDelete)               //删除文章
		ArticleRouter.POST("/createarticlecomment", CreateArticleComment)   //发布对文章的评论
		ArticleRouter.DELETE("/articlecommentdelete", ArticleCommentDelete) //删除文章的评论
		ArticleRouter.GET("/articlesubmited", ArticleSubmited)              //查看某人发布过的文章及其问题和评论
		ArticleRouter.GET("/seearticleinformation", SeeArticleInformation)  //看某一文章信息
		ArticleRouter.GET("/randarticles", RandArticles)                    //返回全部问题
	}
	SearchRouter := r.Group("/api/search") //搜索功能🔍
	{
		SearchRouter.GET("/searchtitle", SearchTitle) //搜索标题
		SearchRouter.GET("/searchall", SearchAll)     //搜索问题回答文章
	}
	PraiseRouter := r.Group("/api/praise") //点赞功能👍
	{
		PraiseRouter.Use(middleware.JWTAuthMiddleware())
		//使用redis缓存
		PraiseRouter.POST("/praisequestion", PraiseQuestion)                           //点赞问题
		PraiseRouter.DELETE("/cancelpraisequestion", CancelPraiseQuestion)             //取消点赞问题
		PraiseRouter.GET("/judgepraisequestion", JudgePraiseQuestion)                  //判断是否点过问题的赞
		PraiseRouter.POST("/praiseanswer", PraiseAnswer)                               //点赞回答
		PraiseRouter.DELETE("/cancelpraiseanswer", CancelPraiseAnswer)                 //取消点赞回答
		PraiseRouter.GET("/judgepraiseanswer", JudgePraiseAnswer)                      //判断是否点过回答的赞
		PraiseRouter.POST("/praisecomment", PraiseComment)                             //点赞回答评论
		PraiseRouter.DELETE("/cancelpraisecomment", CancelPraiseComment)               //取消点赞回答评论
		PraiseRouter.GET("/judgepraisecomment", JudgePraiseComment)                    //判断是否点过回答评论的赞
		PraiseRouter.POST("/praisearticle", PraiseArticle)                             //点赞文章
		PraiseRouter.DELETE("/cancelpraisearticle", CancelPraiseArticle)               //取消点赞文章
		PraiseRouter.GET("/judgepraisearticle", JudgePraiseArticle)                    //判断是否点过问题的赞
		PraiseRouter.POST("/praisearticlecomment", PraiseArticleComment)               //点赞文章评论
		PraiseRouter.DELETE("/cancelpraisearticlecomment", CancelPraiseArticleComment) //取消点赞文章评论
		PraiseRouter.GET("/judgepraisearticlecomment", JudgePraiseArticleComment)      //判断是否点过问题的赞
	}
	r.POST("/api/seepraisequestion", SeePraiseQuestion)             //看问题几个赞
	r.POST("/api/seepraiseanswer", SeePraiseAnswer)                 //看回答几个赞
	r.POST("/api/seepraisecomment", SeePraiseComment)               //看评论几个赞
	r.POST("/api/seepraisearticle", SeePraiseArticle)               //看文章几个赞
	r.POST("/api/seepraisearticlecomment", SeePraiseArticleComment) //看文章评论几个赞
	CollectionRouter := r.Group("/api/collection")                  //收藏功能🗂️
	{
		CollectionRouter.Use(middleware.JWTAuthMiddleware())
		CollectionRouter.POST("/addfavorites", AddFavorites)                      //创建收藏夹
		CollectionRouter.DELETE("/deletefavorites", DeleteFavorites)              //删除收藏夹
		CollectionRouter.POST("/addcollection", AddCollection)                    //收藏回答
		CollectionRouter.POST("/addarticle", AddArticle)                          //收藏文章
		CollectionRouter.DELETE("/deletecollection", DeleteCollection)            //删除收藏回答
		CollectionRouter.DELETE("/deletearticle", DeleteArticle)                  //删除收藏文章
		CollectionRouter.PUT("/modifyfavorites", ModifyFavorites)                 //更改收藏夹属性
		CollectionRouter.GET("/seefavorites", SeeFavorites)                       //看收藏夹内容
		CollectionRouter.GET("/seemyfavorites", SeeMyFavorites)                   //看自己所有的收藏夹
		CollectionRouter.GET("/judgeanswerinfavorites", JudgeAnswerInFavorites)   //判断是否收藏回答
		CollectionRouter.GET("/judgearticleinfavorites", JudgeArticleInFavorites) //判断是否收藏文章
	}
	HotRouter := r.Group("/api/hot") //热榜功能🔥
	{
		HotRouter.GET("/hotanswer", HotAnswer)     //回答热榜
		HotRouter.GET("/hotquestion", HotQuestion) //问题热榜
		HotRouter.GET("/hotarticle", HotArticle)   //文章热榜
	}
	r.POST("/api/addgoods", AddGoods)              //增加商铺
	ShopCenterRouter := r.Group("/api/shopcenter") //商城功能💰
	{
		ShopCenterRouter.Use(middleware.JWTAuthMiddleware())
		ShopCenterRouter.GET("/seegoods", SeeGoods) //看商品
		ShopCenterRouter.POST("/buyvip", BuyVip)    //买VIP
	}
	ChatRouter := r.Group("/api/chat") //聊天功能👩🏼‍❤️‍👨🏻
	{
		ChatRouter.Use(middleware.JWTAuthMiddleware())
		ChatRouter.GET("/seeallunread", service.SeeAllUnread) //看有多少未读信息
		ChatRouter.GET("/seetaunread", service.SeeTAUnread)   //看对于某人有多少未读信息
	}
	r.GET("/api/chatws", service.ChatHandler)
	SubscribeRouter := r.Group("/api/subscribe") //关注通知功能✅
	{
		SubscribeRouter.Use(middleware.JWTAuthMiddleware())
		SubscribeRouter.POST("/subscribepeople", SubscribePeople)               //关注人
		SubscribeRouter.DELETE("/cancelsubscribepeople", CancelSubscribePeople) //取消关注人
		SubscribeRouter.GET("/seepeoplefollowers", SeePeopleFollowers)          //看某人的粉丝数和粉丝id
		SubscribeRouter.GET("/judgefollower", JudgeFollower)                    //判断自己是否关注某人
		SubscribeRouter.GET("/seemysubs", SeeMySubs)                            //看自己的关注
	}
	//用redis实现浏览记录
	r.GET("/api/subsws", service.SubsHandler)
	RecordRouter := r.Group("/api/record")
	{
		RecordRouter.Use(middleware.JWTAuthMiddleware())
		RecordRouter.POST("/addquestionrecord", AddQuestionRecord) //添加一条问题浏览记录
		RecordRouter.POST("/addarticlerecord", AddArticleRecord)   //添加一条文章浏览记录
		RecordRouter.DELETE("/deleteallrecords", DeleteAllRecords) //删除所有记录
		RecordRouter.GET("/seerecord", SeeRecord)                  //看自己的浏览记录
	}
	r.Run(":3920") // 跑在 3920 端口上
}
