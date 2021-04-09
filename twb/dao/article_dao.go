package tdao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"tutil"
)

type Article struct {
	ID         uint32 `db:"idx_article_id"`
	UserID     uint32 `db:"idx_user_id"`
	Title      string `db:"title"`
	Content    string `db:"content"`
	CreateTime uint32 `db:"create_time"`
	UpdateTime uint32 `db:"update_time"`
	Delflag    uint32 `db:"delflag"`
}

func CreateArticle(title string, content string, userID uint32, team uint32) (uint32, error) {
	tableID := (hash32(userID) % 400)

	tutil.Info.Println("CreateArticle", tableID)

	query := fmt.Sprintf("insert into twb_article_tab_%08d(idx_user_id, title, content, create_time, update_time, delflag) values(?,?,?,?,?,?) ", tableID)

	return InsertRow32ID(tableID, query, userID, title, content, time.Now().Unix(), time.Now().Unix(), 0)

}
