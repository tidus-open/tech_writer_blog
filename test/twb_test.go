package main

import (
	//"net/http"
	"fmt"
	"testing"
)

type CreateAccountReq struct {
	Name   string `json:"user_name"`
	Passwd string `json:"passwd"`
}

type CreateAccountRsp struct {
	ID uint32 `json:"user_id"`
}

type CreateTeamReq struct {
	Name   string `json:"team_name"`
	Desc   string `json:"description"`
	UserID uint32 `json:"user_id"`
}

type CommResp struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

type IDResp struct {
	ID uint64 `json:"id"`
}

type CreateArticleReq struct {
	Title   string `json:"title"`
	UserID  uint32 `json:"user_id"`
	Content string `json:"content"`
}

type CreateCommentReq struct {
	ArticleID uint32 `json:"article_id"`
	UserID    uint32 `json:"user_id"`
	Content   string `json:"content"`
}

type UpdateScoreReq struct {
	ArticleID uint32 `json:"article_id"`
	UserID    uint32 `json:"user_id"`
	Score     uint32 `json:"score"`
}

/*
func TestCheckAccount1(t *testing.T) {
	CheckGetApi(t, "https://127.0.0.1:9090/v1/accounts?user_name=delphi&passwd=12345",
		http.StatusOK, CommResp{Code: "OK", Msg: "ok"})
}

func TestCheckAccount2(t *testing.T) {
	CheckGetApi(t, "https://127.0.0.1:9090/v1/accounts?user_name=delphi&passwd=12345",
		http.StatusBadRequest, CommResp{Code: "INVALID_PARAM", Msg: "invalid param"})
}

func TestCreateAccount1(t *testing.T) {
	req := CreateAccountReq{Name: "delphi", Passwd: "12345"}
	CheckPostApi(t, "https://127.0.0.1:9090/v1/accounts", req, http.StatusOK, nil)
}

func TestCreatTeam1(t *testing.T) {
	req := CreateTeamReq{Name: "team1", Desc: "Desc1", UserID: 123456}
	CheckPostApi(t, "https://127.0.0.1:9090/v1/teams", req, http.StatusOK, nil)

}
*/

func iDToTableIDAutoID(ID uint32) (tableID, autoID uint32) {
	tableID = ID & 0x3FF
	autoID = (ID >> 10)
	return tableID, autoID
}

func tableIDAutoIDToID(tableID, autoID uint32) (ID uint32) {
	ID = autoID
	ID <<= 10
	ID += tableID
	return ID
}

func TestArticleLogic1(t *testing.T) {
	aReq := CreateArticleReq{Title: "article1", UserID: 123456, Content: "article1 content"}
	var aRsp = &IDResp{}
	err := PostApi(t, "https://127.0.0.1:9090/v1/articles", aReq, aRsp)
	if err != nil {
		t.Error(err)
		return

	}

	cReq := CreateCommentReq{ArticleID: uint32(aRsp.ID), UserID: 12345, Content: "article 1 comment"}
	cUrl := fmt.Sprintf("https://127.0.0.1:9090/v1/articles/%d/comments", int32(aRsp.ID))
	fmt.Println(cUrl)
	var cRsp = &CommResp{}
	err = PostApi(t, cUrl, cReq, cRsp)
	if err != nil {
		//t.Error(err)
		return
	}

	sReq := UpdateScoreReq{ArticleID: uint32(aRsp.ID), UserID: 12345, Score: 3}
	sUrl := fmt.Sprintf("https://127.0.0.1:9090/v1/articles/%d/score", int32(aRsp.ID))
	fmt.Println(sUrl)
	var sRsp = &CommResp{}
	err = PostApi(t, sUrl, sReq, sRsp)
	if err != nil {
		return
	}

}
