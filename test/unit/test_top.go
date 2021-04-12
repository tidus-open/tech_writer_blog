package ttunit

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
	"ttcomm"
)

func WriteForTop() {
	rand.Seed(time.Now().Unix())
	idx := 0
	for {
		idx++
		articleName := "toparticle" + strconv.Itoa(idx)
		content := "topcontent" + strconv.Itoa(idx)
		_, articleID := CreateArticle(articleName, 12345, content, "")

		ii := (5 + rand.Int()%5)

		for i := 0; i < ii; i++ {
			UpdateScore(articleID, 12345, uint32(1+rand.Int()%4), "")
		}

	}
}

func GetTopList() (*TopListWrap, error) {
	url := fmt.Sprintf("http://127.0.0.1:9090/v1/top")
	fmt.Println(url)

	rsp := &TopListWrap{}
	_, err := ttcomm.GetApi(url, "", rsp)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return rsp, nil

}

func ShowArticle(idx int, articleID uint32, score uint32) {
	url := fmt.Sprintf("http://127.0.0.1:9090/v1/articles/%d", articleID)

	rsp := &Article{}
	_, err := ttcomm.GetApi(url, "", rsp)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d Title: %s score %d content %s updatetime %d\n", idx, rsp.Title, score, rsp.Content, rsp.UpdateTime)

}

func ShowTopList() {

	tpList, err := GetTopList()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("count: %d\n", tpList.Count)
	for i, it := range tpList.List {
		ShowArticle(i, it.ID, it.Score)
	}
}

func ShowTop() {
	for {
		ShowTopList()
		fmt.Println("wait for 5 seconds ........")
		fmt.Printf("\n\n")
		time.Sleep(5 * time.Second)

	}

}
