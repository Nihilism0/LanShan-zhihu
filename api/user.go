package api

import (
	"CSAwork/boot"
	"CSAwork/dao"
	"CSAwork/global"
	"CSAwork/model"
	"CSAwork/utils"
	"CSAwork/utils/middleware"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func Register(c *gin.Context) {
	global.Bucket.Take(1)
	registerform := model.UserRegister{}
	if err := c.ShouldBind(&registerform); err != nil {
		utils.RespFail(c, "verification failed")
		return
	}
	//验证手机号格式
	flag := dao.JudgePhoneNumber(registerform.PhoneNumber)
	if !flag {
		utils.RespFail(c, "phonenumber format error!")
		return
	}
	//验证邮箱格式
	flag = dao.JudgeEmail(registerform.Email)
	if !flag {
		utils.RespFail(c, "email format error!")
		return
	}
	// 验证用户名是否重复
	flag = dao.SelectUser(registerform.UserName)
	if flag {
		// 以 JSON 格式返回信息
		utils.RespFail(c, "user already exist~")
		return
	}
	//验证手机号是否重复
	flag = dao.SelectPhoneNumber(registerform.PhoneNumber)
	if flag {
		// 以 JSON 格式返回信息
		utils.RespFail(c, "phonenumber already exist~")
		return
	}
	//验证邮箱是否重复
	flag = dao.SelectEmail(registerform.Email)
	if flag {
		utils.RespFail(c, "email already exist~")
		return
	}
	//验证验证码
	flag = dao.JudgeCode(strconv.Itoa(int(registerform.PhoneNumber)), registerform.Code)
	if !flag {
		utils.RespFail(c, "Verification Code error~")
		return
	}
	salt := boot.GenValidateCode(6)
	password := boot.MakePassword(registerform.Password, salt)
	TimeNow := time.Now().Format("2006-01-02 15:04:05")
	sqlStr := "insert into users(username,password,created_at,updated_at,phonenumber,email,salt,sign,gender) values (?,?,?,?,?,?,?,?,?)"
	_, _ = global.GlobalDb1.Exec(sqlStr, registerform.UserName, password, TimeNow, TimeNow, registerform.PhoneNumber, registerform.Email, salt, "这是一个为了不出bug而必须设置的超级无聊签名", "女")
	//global.GlobalDb1.
	//	global.GlobalDb1.Model(&model.User{}).Create(&model.User{
	//	Username: registerform.Username,
	//	Password: registerform.Password,
	//})
	// 以 JSON 格式返回信息
	utils.RespSuccess(c, "add user successful")
}

func AccountLogin(c *gin.Context) {
	loginform := model.AccountLogin{}
	if err := c.ShouldBind(&loginform); err != nil {
		utils.RespFail(c, "verification failed")
		return
	}
	// 验证用户名是否存在
	flag := dao.SelectUser(loginform.UserName)
	// 不存在则退出
	if !flag {
		// 以 JSON 格式返回信息
		utils.RespFail(c, "user doesn't exists")
		return
	}

	if dao.JudgePassword(loginform.UserName, loginform.Password) == false {
		utils.RespFail(c, "wrong password")
		return
	}
	// 创建一个我们自己的声明
	claim := model.MyClaims{
		Username: loginform.UserName, // 自定义字段
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(), // 过期时间
			Issuer:    "Joker",                                   // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	tokenString, _ := token.SignedString(middleware.Secret)
	c.JSON(http.StatusOK, gin.H{
		"status":   200,
		"msg":      "success",
		"username": loginform.UserName,
		"data":     gin.H{"token": tokenString},
	})
}

func EmailLogin(c *gin.Context) {
	loginform := model.EmailLogin{}
	if err := c.ShouldBind(&loginform); err != nil {
		utils.RespFail(c, "verification failed")
		return
	}
	flag := dao.JudgeEmail(loginform.Email)
	if !flag {
		utils.RespFail(c, "Email form wrong")
		return
	}
	flag = dao.SelectEmail(loginform.Email)
	if !flag {
		utils.RespFail(c, "Email doesn't exists")
		return
	}
	UserName := dao.GetUsernameFromEmail(loginform.Email)

	if dao.JudgePassword(UserName, loginform.Password) == false {
		utils.RespFail(c, "wrong password")
		return
	}
	// 创建一个我们自己的声明
	claim := model.MyClaims{
		Username: UserName, // 自定义字段
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(), // 过期时间
			Issuer:    "Joker",                                   // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	tokenString, _ := token.SignedString(middleware.Secret)
	c.JSON(http.StatusOK, gin.H{
		"status":   200,
		"msg":      "success",
		"data":     gin.H{"token": tokenString},
		"username": UserName,
	})
}
func PhonePasswordLogin(c *gin.Context) {
	loginform := model.PhonePasswordLogin{}
	if err := c.ShouldBind(&loginform); err != nil {
		utils.RespFail(c, "verification failed")
		return
	}
	flag := dao.JudgePhoneNumber(loginform.PhoneNumber)
	if !flag {
		utils.RespFail(c, "PhoneNumber form wrong")
		return
	}
	flag = dao.SelectPhoneNumber(loginform.PhoneNumber)
	if !flag {
		utils.RespFail(c, "PhoneNumber doesn't exists~")
		return
	}
	UserName := dao.GetUserFromNumber(loginform.PhoneNumber)

	if dao.JudgePassword(UserName, loginform.Password) == false {
		utils.RespFail(c, "wrong password")
		return
	}
	// 创建一个我们自己的声明
	claim := model.MyClaims{
		Username: UserName, // 自定义字段
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(), // 过期时间
			Issuer:    "Joker",                                   // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	tokenString, _ := token.SignedString(middleware.Secret)
	c.JSON(http.StatusOK, gin.H{
		"status":   200,
		"msg":      "success",
		"data":     gin.H{"token": tokenString},
		"username": UserName,
	})
}

func PhoneLogin(c *gin.Context) {
	loginform := model.PhoneLogin{}
	if err := c.ShouldBind(&loginform); err != nil {
		utils.RespFail(c, "verification failed")
		return
	}
	// 验证用户名是否存在
	flag := dao.SelectPhoneNumber(loginform.PhoneNumber)
	// 不存在则退出
	if !flag {
		// 以 JSON 格式返回信息
		utils.RespFail(c, "phonenumber doesn't exists")
		return
	}
	flag = dao.JudgeCode(strconv.Itoa(int(loginform.PhoneNumber)), loginform.Code)
	if !flag {
		utils.RespFail(c, "Verification Code Error!")
		return
	}
	username := dao.GetUserFromNumber(loginform.PhoneNumber)
	// 创建一个我们自己的声明
	claim := model.MyClaims{
		Username: username, // 自定义字段
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(), // 过期时间
			Issuer:    "Joker",                                   // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	tokenString, _ := token.SignedString(middleware.Secret)
	c.JSON(http.StatusOK, gin.H{
		"status":   200,
		"msg":      "success",
		"data":     gin.H{"token": tokenString},
		"username": username,
	})
}

func GetCode(c *gin.Context) {
	global.Bucket.Take(2)
	codenumber := model.CodeNumber{}
	if err := c.ShouldBind(&codenumber); err != nil {
		utils.RespFail(c, "verification failed")
		return
	}
	flag := dao.JudgePhoneNumber(codenumber.PhoneNumber)
	if !flag {
		utils.RespFail(c, "phonenumber format error!")
		return
	}
	boot.Sms(strconv.Itoa(int(codenumber.PhoneNumber)))
	utils.RespSuccess(c, "Send Verification Code Successful!")
}

func InformationModify(c *gin.Context) {
	userinfo := model.UserInfo{}
	if err := c.ShouldBind(&userinfo); err != nil {
		utils.RespFail(c, "verification failed")
		return
	}
	username, _ := c.Get("username")
	if userinfo.UserName != "" {
		if dao.SelectUser(userinfo.UserName) {
			utils.RespFail(c, "The user name is duplicate, please change it~")
			return
		}
	}
	if userinfo.Gender != "" {
		dao.GenderModify(username.(string), userinfo.Gender)
	}
	if userinfo.HeadPhoto != "" {
		dao.HeadPhotoModify(username.(string), userinfo.HeadPhoto)
	}
	if userinfo.Sign != "" {
		dao.SignModify(username.(string), userinfo.Sign)
	}
	if userinfo.UserName != "" {
		dao.UserNameModify(username.(string), userinfo.UserName)
	}
	utils.RespSuccess(c, "Modify information success")
	return
}
func PasswordModify(c *gin.Context) {
	changepassword := model.ChangePassword{}
	if err := c.ShouldBind(&changepassword); err != nil {
		utils.RespFail(c, "verification failed")
		return
	}
	username, _ := c.Get("username")
	UserName := username.(string)
	phonenumber := dao.GetNumberFromUsername(UserName)
	newpassword := changepassword.NewPassword
	code := changepassword.Code
	confirm := changepassword.Confirm
	if dao.JudgePassword(UserName, newpassword) {
		utils.RespFail(c, "The old and new passwords are consistent")
		return
	}
	if newpassword != confirm {
		utils.RespFail(c, "The passwords entered twice are inconsistent")
		return
	}
	if dao.JudgeCode(strconv.Itoa(int(phonenumber)), code) == false {
		utils.RespFail(c, "Code is not true")
		return
	}
	salt := boot.GenValidateCode(6)
	password := boot.MakePassword(newpassword, salt)
	dao.PasswordModify(UserName, password, salt)
	utils.RespSuccess(c, "Change password success")
}
func TokenGetCode(c *gin.Context) {
	username, _ := c.Get("username")
	phonenumber := dao.GetNumberFromUsername(username.(string))
	flag := dao.JudgePhoneNumber(phonenumber)
	if !flag {
		utils.RespFail(c, "phonenumber format error!")
		return
	}
	boot.Sms(strconv.Itoa(int(phonenumber)))
	c.JSON(http.StatusOK, gin.H{
		"status":      200,
		"message":     "Send Verification Code Successful!",
		"phonenumber": phonenumber,
	})
}
func GetUserInformation(c *gin.Context) {
	username, _ := c.Get("username")
	var u model.User
	sqlStr := "select id, created_at, updated_at, username, phonenumber, email, gold,gender, sign,headphoto from users where username=?"
	err := global.GlobalDb1.Get(&u, sqlStr, username)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":      200,
		"information": u,
	})
}
func GetUserHeadPhoto(c *gin.Context) {
	var form struct {
		UserName string `db:"username" form:"username" json:"username" binding:"required"`
	}
	if err := c.ShouldBind(&form); err != nil {
		utils.RespFail(c, "verification failed")
		return
	}
	var u struct {
		HeadPhoto string `db:"headphoto" form:"headphoto" json:"headphoto"`
	}
	sqlStr := "select headphoto from users where username = ?"
	global.GlobalDb1.Get(&u, sqlStr, form.UserName)
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": u.HeadPhoto,
	})
}
