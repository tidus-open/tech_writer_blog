package tdao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"tutil"
)

type Team struct {
	ID         uint32 `db:"idx_team_id"`
	Name       string `db:"idx_team_name"`
	UserID     uint32 `db:"idx_user_id"`
	Desc       string `db:"description"`
	CreateTime uint32 `db:"create_time"`
	UpdateTime uint32 `db:"update_time"`
	MemberCnt  uint32 `db:"member_count"`
}

func GetTeamInfo(id uint32) (*Team, error) {
	tableID, autoID := iDToTableIDAutoID(id)

	tutil.LogInfo("GetTeamInfo id %v", id)

	query := fmt.Sprintf("select idx_team_id, idx_team_name, idx_user_id, description, create_time, update_time, member_count from twb_team_tab_%08d where idx_team_id = ?", tableID)

	var team Team
	err := GetRow(&team, query, autoID)

	return &team, err

}

func CreateTeam(name string, desc string, userID uint32) (uint32, error) {
	tableID := uint32(BKDRHash(name) % 2)

	query := fmt.Sprintf("insert into twb_team_tab_%08d(idx_team_name, idx_user_id, description, create_time, update_time , member_count, delflag) values(?,?,?,?,?,?,?)", tableID)

	return InsertRow32ID(tableID, query, name, userID, desc, time.Now().Unix(), time.Now().Unix(), 0, 0)

}
