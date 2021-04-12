package tlogic

import (
	"tdao"
	"tutil"
)

type Team struct {
	ID         uint32 `json:"team_id"`
	Name       string `json:"team_name"`
	Desc       string `json:"desc"`
	CreateTime uint32 `json:"create_time"`
}

func GetTeamInfo(id uint32) (*Team, error) {
	tutil.LogInfo("GetTeamInfo", id)
	dTeam, err := tdao.GetTeamInfo(id)
	if err != nil {
		return nil, err
	}

	var jt Team
	jt.ID = dTeam.ID
	jt.Name = dTeam.Name
	jt.Desc = dTeam.Desc
	jt.CreateTime = dTeam.CreateTime

	return &jt, nil

}

func CreateTeam(name string, desc string, userID uint32) (id uint32, err error) {
	tutil.LogInfo("CreateTeamInfo")
	id, err = tdao.CreateTeam(name, desc, userID)
	if err != nil {
		tutil.Err.Println(err)
		return 0, err
	}

	return id, nil

}
