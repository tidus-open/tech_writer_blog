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
	tutil.Info.Println("UpdateScore")
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

	tutil.Info.Println("UpdateScoreReq", score)

	if err = tlogic.UpdateScore(score.ArticleID, score.Score); err != nil {
		httpBadRequest(w)
		return
	}

	httpOK(w)

}

func GetScore(w http.ResponseWriter, r *http.Request) {
	tutil.Info.Println("GetScore")

	err := checkToken(r)
	if err != nil {
		httpBadRequest(w)
		return
	}

	articleID, _ := getParamUInt(r.URL.Query(), "article_id")

	tutil.Info.Println("GetScore", articleID)

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
