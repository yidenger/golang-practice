package v1

import (
	"log"
	"golang-practice/blog/pkg/e"
	"github.com/astaxie/beego/validation"
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
)

func GetArticle(c *gin.Context) {
	id, _ := com.StrTo(c.Param("id")).Int()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	var data interface {}

	if !valid.HasErrors() {
		if models.ExistArticleByID(id) {
			data = models.GetArticle(id)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": data,
	})
}

func GetArticles(c *gin.Context) {

}

func AddArticle(c *gin.Context) {

}

func EditArticle(c *gin.Context) {

}

func DeleteArticle(c *gin.Context) {
	
}