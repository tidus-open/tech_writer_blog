package tlogic

import (
	"tdao"
	"tutil"
)

type Score struct {
	Score uint32 `json:"score"`
}

func UpdateScore(articleID uint32, score uint32) error {
	tutil.LogInfo("UpdateScore", articleID, score)

	return tdao.UpdateScore(articleID, score)
	/*
		err := tdao.UpdateScore(articleID, score)
		if err != nil {
			return err
		}

		dbscore, err1 := tdao.GetScore(articleID)
		if err1 != nil {
			return err1
		}

		UpdateTopList(int32(dbscore.TotalScore), articleID)
	*/

	return nil
}

func GetScore(articleID uint32) (*Score, error) {
	tutil.LogInfo("GetScore", articleID)

	dbscore, err := tdao.GetScore(articleID)
	if err != nil {
		return nil, err
	}

	score := &Score{}

	if dbscore.MemberCnt > 0 {
		score.Score = uint32(dbscore.TotalScore / dbscore.MemberCnt)
	}

	return score, nil
}
