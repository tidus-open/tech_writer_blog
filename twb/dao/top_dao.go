package tdao

import (
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
	"time"
	"tutil"
)

type TopList struct {
	Score     uint32
	ArticleID uint32
}

var tpKey = "twb_top_list"

func TAddItem(score int32, articleID uint32) error {
	fmt.Println(redisCli)
	err := redisCli.ZAdd(tpKey,
		redis.Z{Score: float64(score), Member: strconv.Itoa(int(articleID))}).Err()
	if err != nil {
		tutil.Err.Println(err)
		return err
	}

	return nil
}

func TCount() int64 {

	count, err := redisCli.ZCard(tpKey).Result()
	if err != nil {
		tutil.Err.Println(err)
		return 0
	}

	return count

}

func GetTop() ([]TopList, error) {
	op := redis.ZRangeBy{
		Min:    "0",
		Max:    "999999",
		Offset: 0,
		Count:  10,
	}

	vals, err := redisCli.ZRangeByScoreWithScores(tpKey, op).Result()
	if err != nil {
		tutil.Err.Println(err)
		return nil, err
	}

	topList := make([]TopList, len(vals))

	for i, val := range vals {
		topList[i].Score = uint32(val.Score)
		member, _ := val.Member.(string)
		id, _ := strconv.Atoi(member)
		topList[i].ArticleID = uint32(id)
	}

	return topList, nil

}

func TMinScore() int32 {
	op := redis.ZRangeBy{
		Min:    "0",
		Max:    "999999",
		Offset: 0,
		Count:  10,
	}

	vals, err := redisCli.ZRangeByScoreWithScores(tpKey, op).Result()
	if err != nil {
		tutil.Err.Println(err)
		return 0
	}

	for _, val := range vals {
		fmt.Println(val)
	}

	return 0

}

func TDelMinScore() {
	redisCli.ZRemRangeByRank(tpKey, 0, 0)
}

func RefreshTopList() {
	for {
		fmt.Println("RefreshTopList")
		op := redis.ZRangeBy{
			Min:    "0",
			Max:    "999999",
			Offset: 0,
			Count:  100,
		}

		vals, err := redisCli.ZRangeByScoreWithScores(tpKey, op).Result()
		if err != nil {
			tutil.Err.Println(err)
			return
		}

		for _, val := range vals {
			member, _ := val.Member.(string)
			redisCli.ZIncrBy(tpKey, -1, member)
		}
		time.Sleep(8 * time.Second)
	}

}
