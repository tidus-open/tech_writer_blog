package tlogic

import (
	"tdao"
)

type TopList struct {
	Score uint32 `json:"score"`
	ID    uint32 `json:"id"`
}

type TopListWrap struct {
	Count int `json:"count"`
	List  []TopList
}

func GetTop() (*TopListWrap, error) {

	daoList, err := tdao.GetTop()
	if err != nil {
		return nil, err
	}

	var tList TopListWrap
	tList.Count = len(daoList)
	tList.List = make([]TopList, len(daoList))

	for i, val := range daoList {
		tList.List[i].Score = val.Score
		tList.List[i].ID = val.ArticleID
	}

	return &tList, err
}

func UpdateTopList(score int32, articleID uint32) {
	if tdao.TCount() < 10 {
		tdao.TAddItem(score, articleID)
	} else {
		if tdao.TMinScore() < score {
			tdao.TDelMinScore()
			tdao.TAddItem(score, articleID)
		}
	}
}

func RefreshTopList() {
	tdao.RefreshTopList()
}
