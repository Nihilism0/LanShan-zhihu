package dao

import (
	"CSAwork/global"
	"CSAwork/model"
	"fmt"
)

// SelectFavoritesName SelectFavorites 判断收藏夹是否存在
func SelectFavoritesName(username, favoritesname string) bool {
	var u model.Favorites
	sqlStr := "select favoritesname from favorites where username=? and favoritesname=?"
	_ = global.GlobalDb1.Get(&u, sqlStr, username, favoritesname)
	if u.FavoritesName == "" {
		return false
	}
	return true
}
func SelectFavoritesID(username string, id uint) bool {
	var u model.Favorites
	sqlStr := "select username from favorites where id = ?"
	_ = global.GlobalDb1.Get(&u, sqlStr, id)
	if u.UserName == username {
		return true
	}
	return false
}

// AddFavorites 创建收藏夹
func AddFavorites(username, favoritesname string, private uint, describe string) {
	sqlStr := "insert into favorites(username, favoritesname,private,`describe`) values (?,?,?,?)"
	_, err := global.GlobalDb1.Exec(sqlStr, username, favoritesname, private, describe)
	if err != nil {
		fmt.Println(err)
	}
}

// DeleteFavorites 删除收藏夹
func DeleteFavorites(username string, id uint) {
	sqlStr := "delete from favorites where username = ? and id = ?"
	global.GlobalDb1.Exec(sqlStr, username, id)
}

// JudgeFavorites 判断收藏夹是否属于用户
func JudgeFavorites(username string, id uint) bool {
	var u model.Favorites
	sqlStr := "select username from favorites where username=? and id=?"
	_ = global.GlobalDb1.Get(&u, sqlStr, username, id)
	if u.UserName == "" {
		return false
	}
	return true
}

// JudgeFavoritesPrivate 检查是否有查看收藏夹权限
func JudgeFavoritesPrivate(username string, id uint) bool {
	if JudgeFavorites(username, id) {
		return true
	}
	var u model.Favorites
	sqlStr := "select private from favorites where id = ?"
	_ = global.GlobalDb1.Get(&u, sqlStr, id)
	if u.Private == 1 {
		return false
	}
	return true
}

// AddCollection 创建收藏文章
func AddCollection(favorites_id, answer_id uint) {
	sqlStr := "insert into collections(favorites_id, answer_id) values (?,?)"
	global.GlobalDb1.Exec(sqlStr, favorites_id, answer_id)
}
func AddArticleCollection(favorites_id, article_id uint) {
	sqlStr := "insert into collections(favorites_id, article_id) values (?,?)"
	global.GlobalDb1.Exec(sqlStr, favorites_id, article_id)
}

// DeleteCollection 删除收藏文章
func DeleteCollection(favorites_id, answer_id uint) {
	sqlStr := "delete from collections where favorites_id = ? and answer_id = ?"
	global.GlobalDb1.Exec(sqlStr, favorites_id, answer_id)
}
func DeleteArticleCollection(favorites_id, article_id uint) {
	sqlStr := "delete from collections where favorites_id = ? and article_id = ?"
	global.GlobalDb1.Exec(sqlStr, favorites_id, article_id)
}

// JudgeAnswerInFavorites 判断文章是否在里面收藏夹里
func JudgeAnswerInFavorites(favorites_id, answer_id uint) bool {
	var u model.Collection
	sqlStr := "select favorites_id from collections where favorites_id= ? and answer_id= ? "
	global.GlobalDb1.Get(&u, sqlStr, favorites_id, answer_id)
	if u.Favorites_id == 0 {
		return false
	}
	return true
}
func JudgeArticleInFavorites(favorites_id, article_id uint) bool {
	var u model.Collection
	sqlStr := "select favorites_id from collections where favorites_id= ? and article_id= ? "
	global.GlobalDb1.Get(&u, sqlStr, favorites_id, article_id)
	if u.Favorites_id == 0 {
		return false
	}
	return true
}

func FavoritesNameModify(id uint, name string) {
	sqlStr := "update favorites set favoritesname = ? where id = ?"
	_, _ = global.GlobalDb1.Exec(sqlStr, name, id)
}

func FavoritesPrivateModify(id, private uint, describe string) {
	sqlStr := "update favorites set private = ? where id = ?"
	_, _ = global.GlobalDb1.Exec(sqlStr, private, id)
	if describe != "" {
		sqlStr = "update favorites set `describe` = ? where id = ?"
		_, _ = global.GlobalDb1.Exec(sqlStr, describe, id)
	}
}

func FindFavorites(id uint) []model.RealCollection {
	var collections []model.RealCollection
	sqlStr := "select * from collections where favorites_id = ? "
	err := global.GlobalDb1.Select(&collections, sqlStr, id)
	if err != nil {
		fmt.Println(err)
	}
	return collections
}

func FindMyFavorites(username string) []model.Favorites {
	var favorites []model.Favorites
	sqlStr := "select id, username, favoritesname, private, `describe` from favorites where username = ?"
	err := global.GlobalDb1.Select(&favorites, sqlStr, username)
	if err != nil {
		println(username, "   err!!!:", err)
	}
	return favorites
}
func CancelAnswerInFavorites(answer_id uint) {
	sqlStr := "delete from collections where answer_id = ?"
	_, _ = global.GlobalDb1.Exec(sqlStr, answer_id)
}
func CancelArticleInFavorites(article_id uint) {
	sqlStr := "delete from collections where article_id = ?"
	_, _ = global.GlobalDb1.Exec(sqlStr, article_id)
}
