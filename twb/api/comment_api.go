package tapi

import (
	"net/http"
	"tlogic"
	"tutil"
)

type CreateCommentReq struct {
	ArticleID uint32 `json:"article_id"`
	UserID    uint32 `json:"user_id"`
	Content   string `json:"content"`
}

func CreateComment(w http.ResponseWriter, r *http.Request) {
	tutil.LogInfo("CreateComment")

	var comment = &CreateCommentReq{}
	err := checkToken(r)
	if err != nil {
		httpBadRequest(w)
		return
	}

	err = GetJsonParam(comment, r)
	if err != nil {
		httpBadRequest(w)
		return
	}

	tutil.LogInfo("CreateComment", comment)

	if _, err = tlogic.CreateComment(comment.ArticleID, comment.UserID, comment.Content); err != nil {
		httpBadRequest(w)
		return
	}

	httpOK(w)

}

func GetComment(w http.ResponseWriter, r *http.Request) {
	tutil.LogInfo("GetComment")

	articleID, _ := getIdFromURL(r.URL.Path, 3)

	err := checkToken(r)
	if err != nil {
		httpBadRequest(w)
		return
	}

	tutil.LogInfo("GetComment", articleID)

	comments, err := tlogic.GetComment(uint32(articleID))
	if err == tutil.ErrNotFound {
		httpBadRequest(w)
		return
	}

	if err == tutil.InternalErr {
		httpInterErr(w)
		return
	}

	httpResp(w, comments)

}
