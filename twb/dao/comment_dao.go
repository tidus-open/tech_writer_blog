package tdao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"tutil"
)

type Comment struct {
	ArticleCreateID uint64 `db:"idx_article_create"`
	UserID          uint32 `db:"user_id"`
	CreateTime      uint32 `db:"create_time"`
	Content         string `db:"content"`
}

func CreateComment(articleID uint32, userID uint32, content string) (uint64, error) {
	tableID := (hash32(articleID) % 1000)

	tutil.Info.Println("CreateComment", tableID)

	query := fmt.Sprintf("insert into twb_comment_tab_%08d(idx_article_create, user_id, create_time, content, delflag) values(?,?,?,?,?)", tableID)

	articleCreate := uint64(articleID)
	articleCreate <<= 32
	articleCreate += uint64(time.Now().Unix())

	return InsertRow64ID(tableID, query, articleCreate, userID, time.Now().Unix(), content, 0)

}

func GetComment(articleID uint32) ([]Comment, error) {
	tableID := (hash32(articleID) % 1000)

	tutil.Info.Println("GetComment", tableID)

	articleCreate := uint64(articleID)
	articleCreate <<= 32

	start := articleCreate
	end := (articleCreate | 0xffffffff)

	query := fmt.Sprintf("select idx_article_create, user_id, create_time, content from twb_comment_tab_%08d where  idx_article_create > ? and idx_article_create < ?")

	var comment []Comment

	err := GetRows(comment, query, start, end)

	return comment, err

}
