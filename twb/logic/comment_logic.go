package tlogic

import (
	"tdao"
	"tutil"
)

type Comment struct {
	UserID     uint32 `json:"user_id"`
	Content    string `json:"content"`
	CreateTime uint32 `json:"create_time"`
}

func CreateComment(articleID uint32, userID uint32, content string) (uint64, error) {
	tutil.Info.Println("CreateComment", articleID, userID, content)

	return tdao.CreateComment(articleID, userID, content)

}

func GetComment(articleID uint32) ([]Comment, error) {
	tutil.Info.Println("GetComment", articleID)

	dbcomment, err := tdao.GetComment(articleID)
	if err != nil {
		return nil, err
	}

	comment := make([]Comment, len(dbcomment))

	for i, v := range dbcomment {
		comment[i].UserID = v.UserID
		comment[i].Content = v.Content
		comment[i].CreateTime = v.CreateTime
	}

	return comment, nil

}
