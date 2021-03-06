package tdao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"tutil"
)

type Score struct {
	ArticleID  uint32 `db:"idx_article_id"`
	TotalScore uint32 `db:"total_score"`
	MemberCnt  uint32 `db:"member_cnt"`
}

func CreateScore(articleID uint32) (uint32, error) {
	tableID := (hash32(articleID) % 1000)

	tutil.LogInfo("CreateScore", articleID, tableID)

	query := fmt.Sprintf("insert into twb_article_score_tab_%08d(idx_article_id, total_score, member_cnt) values(?,?,?)", tableID)

	return InsertRow32ID(tableID, query, articleID, 0, 0)

}

func UpdateScore(articleID uint32, score uint32) error {
	tableID := (hash32(articleID) % 1000)

	tutil.LogInfo("UpdateScore", articleID, tableID)

	query := fmt.Sprintf("update twb_article_score_tab_%08d set total_score=total_score+? , member_cnt=member_cnt+1 where idx_article_id = ?", tableID)

	return UpdateRow(tableID, query, score, articleID)
}

func GetScore(articleID uint32) (*Score, error) {
	tableID := (hash32(articleID) % 1000)

	tutil.LogInfo("GetScore", tableID, articleID)

	query := fmt.Sprintf("select idx_article_id, total_score, member_cnt from twb_article_score_tab_%08d where idx_article_id = ?", tableID)

	var score Score
	err := GetRow(&score, query, articleID)

	tutil.Info.Println(score)

	return &score, err

}
