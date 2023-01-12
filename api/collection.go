package api

import (
	"CSAwork/dao"
	"CSAwork/global"
	"CSAwork/model"
	"CSAwork/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddFavorites(c *gin.Context) {
	global.Bucket.Take(1)
	favoritesform := model.Favorites{}
	if err := c.ShouldBind(&favoritesform); err != nil {
		fmt.Println(err)
		utils.RespFail(c, "submit form wrong")
		return
	}
	username, _ := c.Get("username")
	favoritesname := favoritesform.FavoritesName
	private := favoritesform.Private
	describe := favoritesform.Describe
	flag := dao.SelectFavoritesName(username.(string), favoritesname)
	if flag == true {
		utils.RespFail(c, "You already behave this favorites")
		return
	}
	dao.AddFavorites(username.(string), favoritesname, private, describe)
	utils.RespSuccess(c, "Add favorites success")
}
func DeleteFavorites(c *gin.Context) {
	favoritesform := model.FavoritesDelete{}
	if err := c.ShouldBind(&favoritesform); err != nil {
		utils.RespFail(c, "submit form wrong")
		return
	}
	username, _ := c.Get("username")
	favoritesid := favoritesform.ID
	flag := dao.JudgeFavorites(username.(string), favoritesid)
	if flag == false {
		utils.RespFail(c, "favorites is not belong to you")
		return
	}
	dao.DeleteFavorites(username.(string), favoritesid)
	utils.RespSuccess(c, "Delete favorites success")
}
func AddCollection(c *gin.Context) {
	collectionform := model.Collection{}
	if err := c.ShouldBind(&collectionform); err != nil {
		utils.RespFail(c, "submit form wrong")
		return
	}
	username, _ := c.Get("username")
	favoritesid := collectionform.Favorites_id
	answerid := collectionform.Answer_id
	flag := dao.JudgeFavorites(username.(string), favoritesid)
	if flag == false {
		utils.RespFail(c, "This favorites not belong to you")
		return
	}
	flag = dao.JudgeAnswerInFavorites(favoritesid, answerid)
	if flag == true {
		utils.RespFail(c, "You already add this answer")
		return
	}
	dao.AddCollection(favoritesid, answerid)
	utils.RespSuccess(c, "Add answer success")
}

func AddArticle(c *gin.Context) {
	collectionform := model.ArticleCollection{}
	if err := c.ShouldBind(&collectionform); err != nil {
		utils.RespFail(c, "submit form wrong")
		return
	}
	username, _ := c.Get("username")
	favoritesid := collectionform.Favorites_id
	articleid := collectionform.Article_id
	flag := dao.JudgeFavorites(username.(string), favoritesid)
	if flag == false {
		utils.RespFail(c, "This favorites not belong to you")
		return
	}
	flag = dao.JudgeArticleInFavorites(favoritesid, articleid)
	if flag == true {
		utils.RespFail(c, "You already add this article")
		return
	}
	dao.AddArticleCollection(favoritesid, articleid)
	utils.RespSuccess(c, "Add article success")
}

func DeleteCollection(c *gin.Context) {
	collectionform := model.Collection{}
	if err := c.ShouldBind(&collectionform); err != nil {
		utils.RespFail(c, "submit form wrong")
		return
	}
	username, _ := c.Get("username")
	favoritesid := collectionform.Favorites_id
	answerid := collectionform.Answer_id
	flag := dao.JudgeFavorites(username.(string), favoritesid)
	if flag == false {
		utils.RespFail(c, "This favorites not belong to you")
		return
	}
	flag = dao.JudgeAnswerInFavorites(favoritesid, answerid)
	if flag == false {
		utils.RespFail(c, "Answer is not in favorites")
		return
	}
	dao.DeleteCollection(favoritesid, answerid)
	utils.RespSuccess(c, "Delete collection success")
}
func DeleteArticle(c *gin.Context) {
	collectionform := model.ArticleCollection{}
	if err := c.ShouldBind(&collectionform); err != nil {
		utils.RespFail(c, "submit form wrong")
		return
	}
	username, _ := c.Get("username")
	favoritesid := collectionform.Favorites_id
	articleid := collectionform.Article_id
	flag := dao.JudgeFavorites(username.(string), favoritesid)
	if flag == false {
		utils.RespFail(c, "This favorites not belong to you")
		return
	}
	flag = dao.JudgeArticleInFavorites(favoritesid, articleid)
	if flag == false {
		utils.RespFail(c, "Article is not in favorites")
		return
	}
	dao.DeleteArticleCollection(favoritesid, articleid)
	utils.RespSuccess(c, "Delete collection success")
}
func ModifyFavorites(c *gin.Context) {
	favoritesmodify := model.FavoritesModify{}
	if err := c.ShouldBind(&favoritesmodify); err != nil {
		utils.RespFail(c, "submit form wrong")
		return
	}
	username, _ := c.Get("username")
	id := favoritesmodify.ID
	favoritesname := favoritesmodify.FavoritesName
	private := favoritesmodify.Private
	describe := favoritesmodify.Describe
	if dao.SelectFavoritesID(username.(string), id) == false {
		utils.RespFail(c, "This favorites is not belong to you")
		return
	}
	flag := dao.SelectFavoritesName(username.(string), favoritesname)
	if flag == true {
		utils.RespFail(c, "You already behave this favorites name")
		return
	}
	dao.FavoritesPrivateModify(id, private, describe)
	if favoritesname != "" {
		dao.FavoritesNameModify(id, favoritesname)
	}
	utils.RespSuccess(c, "Modify favorites success")
}

func SeeFavorites(c *gin.Context) {

	SeeFavoritesForm := model.FavoritesID{}
	if err := c.ShouldBind(&SeeFavoritesForm); err != nil {
		utils.RespFail(c, "submit form wrong")
		return
	}
	username, _ := c.Get("username")
	id := SeeFavoritesForm.ID
	if dao.JudgeFavoritesPrivate(username.(string), id) == false {
		utils.RespFail(c, "You do not have authority to see")
		return
	}
	collections := dao.FindFavorites(id)
	c.JSON(http.StatusOK, gin.H{
		"status":      200,
		"collections": collections,
	})
}
func SeeMyFavorites(c *gin.Context) {
	username, _ := c.Get("username")
	favorites := dao.FindMyFavorites(username.(string))
	c.JSON(http.StatusOK, gin.H{
		"status":    200,
		"favorites": favorites,
	})
}
func JudgeAnswerInFavorites(c *gin.Context) {
	SeeFavoritesForm := model.FavoritesID{}
	username, _ := c.Get("username")
	if err := c.ShouldBind(&SeeFavoritesForm); err != nil {
		utils.RespFail(c, "submit form wrong")
		return
	}
	favorites := dao.FindMyFavorites(username.(string))
	flag := 0
	var ID []uint
	for _, v := range favorites {
		if dao.JudgeAnswerInFavorites(v.ID, SeeFavoritesForm.ID) {
			ID = append(ID, v.ID)
			flag = 1
		}
	}
	if flag == 0 {
		utils.RespFail(c, "没收藏")
		return
	} else if flag == 1 {
		c.JSON(http.StatusOK, gin.H{
			"status":       200,
			"favorites_id": ID,
			"message":      "收藏过了",
		})
		return
	}
}
func JudgeArticleInFavorites(c *gin.Context) {
	SeeFavoritesForm := model.FavoritesID{}
	username, _ := c.Get("username")
	if err := c.ShouldBind(&SeeFavoritesForm); err != nil {
		utils.RespFail(c, "submit form wrong")
		return
	}
	favorites := dao.FindMyFavorites(username.(string))
	flag := 0
	var ID []uint
	for _, v := range favorites {
		if dao.JudgeArticleInFavorites(v.ID, SeeFavoritesForm.ID) {
			ID = append(ID, v.ID)
			flag = 1
		}
	}
	if flag == 0 {
		utils.RespFail(c, "没收藏")
		return
	} else if flag == 1 {
		c.JSON(http.StatusOK, gin.H{
			"status":     200,
			"article_id": ID,
			"message":    "收藏过了",
		})
		return
	}
}
