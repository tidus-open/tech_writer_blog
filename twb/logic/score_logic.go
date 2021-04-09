package tlogic

import (
	"tdao"
	"tutil"
)

type Score struct {
	Score uint32 `json:"score"`
}

func UpdateScore(articleID uint32, score uint32) error {
	tutil.Info.Println("UpdateScore", articleID, score)

	return tdao.UpdateScore(articleID, score)

}

func GetScore(articleID uint32) (*Score, error) {
	tutil.Info.Println("GetScore", articleID)

	dbscore, err := tdao.GetScore(articleID)
	if err != nil {
		return nil, err
	}

	score := &Score{}

	score.Score = uint32(dbscore.TotalScore / dbscore.MemberCnt)

	return score, nil
}
