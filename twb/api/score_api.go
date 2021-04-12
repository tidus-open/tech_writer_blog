package tapi

import (
	"net/http"
	"tlogic"
	"tutil"
)

type UpdateScoreReq struct {
	ArticleID uint32 `json:"article_id"`
	UserID    uint32 `json:"user_id"`
	Score     uint32 `json:"score"`
}

func UpdateScore(w http.ResponseWriter, r *http.Request) {
	tutil.LogInfo("UpdateScore")
	var score = &UpdateScoreReq{}

	err := checkToken(r)
	if err != nil {
		httpBadRequest(w)
		return
	}

	err = GetJsonParam(score, r)
	if err != nil {
		httpBadRequest(w)
		return
	}

	tutil.LogInfo("UpdateScoreReq", score)

	if err = tlogic.UpdateScore(score.ArticleID, score.Score); err != nil {
		httpBadRequest(w)
		return
	}

	httpOK(w)

}

func GetScore(w http.ResponseWriter, r *http.Request) {
	tutil.LogInfo("GetScore")

	articleID, _ := getIdFromURL(r.URL.Path, 3)

	err := checkToken(r)
	if err != nil {
		httpBadRequest(w)
		return
	}

	tutil.LogInfo("GetScore", articleID)

	score, err := tlogic.GetScore(uint32(articleID))
	if err == tutil.ErrNotFound {
		httpBadRequest(w)
		return
	}

	if err == tutil.InternalErr {
		httpInterErr(w)
		return
	}

	httpResp(w, score)

}
