package tapi

import (
	"net/http"
	"tlogic"
	"tutil"
)

type Article struct {
	Title   string `json:"title"`
	UserID  uint32 `json:"user_id"`
	Content string `json:"content"`
}

type CreateArticleRsp struct {
	ID uint32 `json:"id"`
}

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	tutil.LogInfo("CreateArticle")

	var at = &Article{}
	err := checkToken(r)
	if err != nil {
		httpBadRequest(w)
		return
	}

	err = GetJsonParam(at, r)
	if err != nil {
		httpBadRequest(w)
		return
	}

	tutil.LogInfo("CreateArticle", at.Title, at.Content, at.UserID)

	var articleID uint32
	if articleID, err = tlogic.CreateArticle(at.Title, at.Content, at.UserID); err != nil {
		httpBadRequest(w)
		return
	}

	tutil.LogInfo("CreateArticle ID", articleID)

	rsp := CreateArticleRsp{ID: articleID}
	httpResp(w, rsp)

}

func GetArticle(w http.ResponseWriter, r *http.Request) {
	tutil.LogInfo("GetArticle")

	articleID, _ := getIdFromURL(r.URL.Path, 3)

	err := checkToken(r)
	if err != nil {
		httpBadRequest(w)
		return
	}

	tutil.LogInfo("GetArticle", articleID)

	article, err := tlogic.GetArticle(uint32(articleID))
	if err == tutil.ErrNotFound {
		httpBadRequest(w)
		return
	}

	if err == tutil.InternalErr {
		httpInterErr(w)
		return
	}

	httpResp(w, article)

}
