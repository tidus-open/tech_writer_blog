package tapi

import (
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
	tutil.Info.Println("GetTeamInfo")

	err := checkToken(r)
	if err != nil {
		httpBadRequest(w)
		return
	}

	teamID, _ := getParamUInt(r.URL.Query(), "team_id")

	tutil.Info.Println("GetTeamInfo", teamID)

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
		httpBadRequest(w)
		return
	}

	err = GetJsonParam(req, r)
	if err != nil {
		httpBadRequest(w)
		return
	}

	tutil.Info.Println("CreateTeam", req.Name, req.Desc, req.UserID)

	var autoID uint32
	if autoID, err = tlogic.CreateTeam(req.Name, req.Desc, req.UserID); err != nil {
		httpBadRequest(w)
		return
	}

	rsp := CreateTeamRsp{ID: autoID}
	httpResp(w, rsp)

}
