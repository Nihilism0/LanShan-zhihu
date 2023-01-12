package dao

import (
	"CSAwork/boot"
	"CSAwork/global"
	"CSAwork/model"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// 若没有这个用户返回 false，反之返回 true
func SelectUser(username string) bool {
	var u model.UserInfo
	sqlStr := "select username from users where username=?"
	_ = global.GlobalDb1.Get(&u, sqlStr, username)
	//global.GlobalDb1.Model(&model.User{}).Where("username = ?", username).Find(&u)
	if u.UserName == "" {
		return false
	} else {
		return true
	}
}
func SelectPhoneNumber(phonenumber uint) bool {
	var u model.PhoneLogin
	sqlStr := "select phonenumber from users where phonenumber=?"
	_ = global.GlobalDb1.Get(&u, sqlStr, phonenumber)
	if u.PhoneNumber == 0 {
		return false
	} else {
		return true
	}
}
func SelectEmail(email string) bool {
	var u model.EmailLogin
	sqlStr := "select email from users where email=?"
	_ = global.GlobalDb1.Get(&u, sqlStr, email)
	if u.Email == "" {
		return false
	} else {
		return true
	}
}
func JudgePhoneNumber(phonenumber uint) bool {
	num := strconv.Itoa(int(phonenumber))
	if !strings.HasPrefix(num, "1") {
		return false
	}
	if len([]rune(num)) != 11 {
		return false
	}
	return true
}
func JudgeEmail(email string) bool {
	if strings.HasSuffix(email, "@qq.com") {
		return true
	}
	return false
}
func JudgePassword(username, password string) bool {
	var u model.GetPassword
	sqlStr := "select password,salt from users where username=?"
	global.GlobalDb1.Get(&u, sqlStr, username)
	//global.GlobalDb1.Model(&model.User{}).Where("username = ?", username).Find(&u)
	flag := boot.ValidPassword(password, u.Salt, u.Password)
	return flag
}
func GetUsernameFromEmail(email string) string {
	var u model.UserInfo
	sqlStr := "select username from users where email=?"
	// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	global.GlobalDb1.Get(&u, sqlStr, email)
	//global.GlobalDb1.Model(&model.User{}).Where("username = ?", username).Find(&u)
	return u.UserName
}
func JudgeCode(phonenumber string, code string) bool {
	val, _ := global.RedisDb.Get(phonenumber).Result()
	if code == val {
		return true
	}
	return false
}

func GetUserFromNumber(phonenumber uint) string {
	var u model.UserInfo
	sqlStr := "select username from users where phonenumber=?"
	global.GlobalDb1.Get(&u, sqlStr, phonenumber)
	return u.UserName
}

func GenderModify(username, gender string) {
	sqlStr := "update users set gender=? where username = ?"
	global.GlobalDb1.Exec(sqlStr, gender, username)
	timenow := time.Now().Format("2006-01-02 15:04:05")
	sqlStr = "update users set updated_at=? where username = ?"
	global.GlobalDb1.Exec(sqlStr, timenow, username)
}

func SignModify(username, sign string) {
	sqlStr := "update users set sign=? where username = ?"
	global.GlobalDb1.Exec(sqlStr, sign, username)
	timenow := time.Now().Format("2006-01-02 15:04:05")
	sqlStr = "update users set updated_at=? where username = ?"
	global.GlobalDb1.Exec(sqlStr, timenow, username)
}
func HeadPhotoModify(username, headphoto string) {
	sqlStr := "update users set headphoto=? where username = ?"
	global.GlobalDb1.Exec(sqlStr, headphoto, username)
	timenow := time.Now().Format("2006-01-02 15:04:05")
	sqlStr = "update users set updated_at=? where username = ?"
	global.GlobalDb1.Exec(sqlStr, timenow, username)
}
func UserNameModify(oldusername, newusername string) {
	timenow := time.Now().Format("2006-01-02 15:04:05")
	sqlStr := "update users set updated_at=? where username = ?"
	_, err := global.GlobalDb1.Exec(sqlStr, timenow, oldusername)
	if err != nil {
		fmt.Println(err)
	}
	sqlStr = "update users set username=?  where username = ?"
	_, err = global.GlobalDb1.Exec(sqlStr, newusername, oldusername)
	if err != nil {
		fmt.Println(err)
	}
}

func GetNumberFromUsername(username string) uint {
	var u model.UserRegister
	sqlStr := "select phonenumber from users where username=?"
	global.GlobalDb1.Get(&u, sqlStr, username)
	return u.PhoneNumber
}
func PasswordModify(username, newpassword, salt string) {
	timenow := time.Now().Format("2006-01-02 15:04:05")
	sqlStr := "update users set updated_at=? where username = ?"
	_, err := global.GlobalDb1.Exec(sqlStr, timenow, username)
	if err != nil {
		fmt.Println(err)
	}
	sqlStr = "update users set password=? where username = ?"
	_, err = global.GlobalDb1.Exec(sqlStr, newpassword, username)
	if err != nil {
		fmt.Println(err)
	}
	sqlStr = "update users set salt=? where username = ?"
	_, err = global.GlobalDb1.Exec(sqlStr, salt, username)
	if err != nil {
		fmt.Println(err)
	}
}
func GetIdFromUsername(username string) uint {
	var u model.User
	sqlStr := "select id from users where username= ?"
	global.GlobalDb1.Get(&u, sqlStr, username)
	return u.ID
}
