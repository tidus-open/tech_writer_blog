package tapi

import (
	//"fmt"
	"net/http"
	"tlogic"
	"tutil"
)

type CreateTeamReq struct {
	Name   string `json:"team_name"`
	Desc   string `json:"description"`
	UserID uint32 `json:"user_id"`
}

type CreateTeamRsp struct {
	ID uint32 `json:"id"`
}

func GetTeamInfo(w http.ResponseWriter, r *http.Request) {
	tutil.LogInfo("GetTeamInfo")
	//fmt.Println(r.URL.Path)

	teamID, _ := getIdFromURL(r.URL.Path, 3)

	err := checkToken(r)
	if err != nil {
		httpBadRequest(w)
		return
	}

	tutil.LogInfo("GetTeamInfo", teamID)

	tinfo, err := tlogic.GetTeamInfo(uint32(teamID))
	if err == tutil.ErrNotFound {
		httpBadRequest(w)
		return
	}

	if err == tutil.InternalErr {
		httpInterErr(w)
		return
	}

	httpResp(w, tinfo)

}

func CreateTeam(w http.ResponseWriter, r *http.Request) {
	var req = &CreateTeamReq{}

	err := checkToken(r)
	if err != nil {
		httpUnauthorized(w)
		return
	}

	err = GetJsonParam(req, r)
	if err != nil {
		httpBadRequest(w)
		return
	}

	tutil.LogInfo("CreateTeam", req.Name, req.Desc, req.UserID)

	var autoID uint32
	if autoID, err = tlogic.CreateTeam(req.Name, req.Desc, req.UserID); err != nil {
		httpBadRequest(w)
		return
	}

	rsp := CreateTeamRsp{ID: autoID}
	httpResp(w, rsp)

}
