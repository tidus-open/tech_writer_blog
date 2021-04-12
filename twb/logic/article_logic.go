package tlogic

import (
	"tdao"
	"tutil"
)

type Article struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	UserID     uint32 `json:"user_id"`
	UpdateTime uint32 `json:"update_time"`
}

func CreateArticle(title string, content string, userID uint32) (uint32, error) {
	tutil.LogInfo("CreateArticle", title, content, userID)

	articleID, err := tdao.CreateArticle(title, content, userID, 0)
	if err != nil {
		return 0, nil
	}

	tutil.LogInfo("CreateArticle ID", articleID)

	_, err = tdao.CreateScore(articleID)
	if err != nil {
		return 0, nil
	}

	return articleID, nil

}

func GetArticle(articleID uint32) (*Article, error) {
	tutil.LogInfo("GetArticle", articleID)
	dbarticle, err := tdao.GetArticle(articleID)
	if err != nil {
		return nil, err
	}

	var article Article

	article.Title = dbarticle.Title
	article.Content = dbarticle.Content
	article.UserID = dbarticle.UserID
	article.UpdateTime = dbarticle.UpdateTime

	return &article, nil

}
