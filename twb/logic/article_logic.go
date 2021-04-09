package tlogic

import (
	"tdao"
	"tutil"
)

func CreateArticle(title string, content string, userID uint32) (uint32, error) {
	tutil.Info.Println("CreateArticle", title, content, userID)

	articleID, err := tdao.CreateArticle(title, content, userID, 0)
	if err != nil {
		return 0, nil
	}

	tutil.Info.Println("CreateArticle ID", articleID)

	_, err = tdao.CreateScore(articleID)
	if err != nil {
		return 0, nil
	}

	return articleID, nil

}
