package ttunit

import (
	"fmt"
	//"net/http"
	"testing"
)

/*
func TestCheckAccount1(t *testing.T) {
	CheckGetApi("https://127.0.0.1:9090/v1/accounts?user_name=delphi&passwd=12345",
		http.StatusOK, CommResp{Code: "OK", Msg: "ok"})
}

func TestCheckAccount2(t *testing.T) {
	CheckGetApi("https://127.0.0.1:9090/v1/accounts?user_name=delphi&passwd=12345",
		http.StatusBadRequest, CommResp{Code: "INVALID_PARAM", Msg: "invalid param"})
}

func TestCreateAccount1(t *testing.T) {
	req := CreateAccountReq{Name: "delphi", Passwd: "12345"}
	CheckPostApi("https://127.0.0.1:9090/v1/accounts", req, http.StatusOK, nil)
}

func TestCreatTeam1(t *testing.T) {
	req := CreateTeamReq{Name: "team1", Desc: "Desc1", UserID: 123456}
	CheckPostApi("https://127.0.0.1:9090/v1/teams", req, http.StatusOK, nil)

}
*/

/*
func TestWriteLogic(t *testing.T) {
	uReq := CreateAccountReq{Name: "delphi", Passwd: "12345"}
	var uRsp = &IDResp{}
	err := PostApi("https://127.0.0.1:9090/v1/accounts", uReq, uRsp)
	if err != nil {
		t.Error(err)
		return
	}

	userID := uint32(uRsp.ID)

	fmt.Println("userID", userID)

	aReq := CreateArticleReq{Title: "article1", UserID: userID, Content: "article1 content"}
	var aRsp = &IDResp{}
	err = PostApi("https://127.0.0.1:9090/v1/articles", aReq, aRsp)
	if err != nil {
		t.Error(err)
		return

	}

	articleID := uint32(aRsp.ID)
	fmt.Println("articleID", articleID)

	cReq := CreateCommentReq{ArticleID: uint32(aRsp.ID), UserID: userID, Content: "article 1 comment"}
	cUrl := fmt.Sprintf("https://127.0.0.1:9090/v1/articles/%d/comments", articleID)
	fmt.Println(cUrl)
	var cRsp = &CommResp{}
	err = PostApi(cUrl, cReq, cRsp)
	if err != nil {
		t.Error(err)
		return
	}

	sReq := UpdateScoreReq{ArticleID: uint32(aRsp.ID), UserID: 12345, Score: 3}
	sUrl := fmt.Sprintf("https://127.0.0.1:9090/v1/articles/%d/score", articleID)
	fmt.Println(sUrl)
	var sRsp = &CommResp{}
	err = PostApi(sUrl, sReq, sRsp)
	if err != nil {
		t.Error(err)
		return
	}

	tReq := CreateTeamReq{Name: "team1", Desc: "Desc1", UserID: userID}
	var tRsp = &CommResp{}
	err = PostApi("https://127.0.0.1:9090/v1/teams", tReq, tRsp)
	if err != nil {
		t.Error(err)
		return
	}

}
*/

func TestReadLogic(t *testing.T) {

	name := "java"
	passwd := "java123"

	aUrl := fmt.Sprintf("http://127.0.0.1:9090/v1/accounts?user_name=%s&passwd=%s",
		name, passwd)

	aRsp := &CheckAccoutRsp{}
	err := GetApi(aUrl, aRsp)
	if err != nil {
		t.Error(err)
		return
	}

}
