package ttunit

import (
	"fmt"
	"ttcomm"
)

type CreateAccountReq struct {
	Name   string `json:"user_name"`
	Passwd string `json:"passwd"`
}

type CreateAccountRsp struct {
	ID uint32 `json:"user_id"`
}

type CreateTeamReq struct {
	Name   string `json:"team_name"`
	Desc   string `json:"description"`
	UserID uint32 `json:"user_id"`
}

type IDResp struct {
	ID uint64 `json:"id"`
}

type CreateArticleReq struct {
	Title   string `json:"title"`
	UserID  uint32 `json:"user_id"`
	Content string `json:"content"`
}

type CreateCommentReq struct {
	ArticleID uint32 `json:"article_id"`
	UserID    uint32 `json:"user_id"`
	Content   string `json:"content"`
}

type UpdateScoreReq struct {
	ArticleID uint32 `json:"article_id"`
	UserID    uint32 `json:"user_id"`
	Score     uint32 `json:"score"`
}

type CheckAccoutRsp struct {
	Token string `json:"token"`
}

type Team struct {
	ID         uint32 `json:"team_id"`
	Name       string `json:"team_name"`
	Desc       string `json:"desc"`
	CreateTime uint32 `json:"create_time"`
}

type Article struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	UserID     uint32 `json:"user_id"`
	UpdateTime uint32 `json:"update_time"`
}

type Comment struct {
	UserID     uint32 `json:"user_id"`
	Content    string `json:"content"`
	CreateTime uint32 `json:"create_time"`
}

type Score struct {
	Score uint32 `json:"score"`
}

type TopList struct {
	Score uint32 `json:"score"`
	ID    uint32 `json:"id"`
}

type TopListWrap struct {
	Count int `json:"count"`
	List  []TopList
}

func BKDRHash(str string) (hash uint32) {
	seed := uint32(131) // 31 131 1313 13131 131313 etc..
	hash = uint32(0)
	for i := 0; i < len(str); i++ {
		hash = (hash * seed) + uint32(str[i])
	}
	return hash
}

func hash32(a uint32) (b uint32) {
	a = (a ^ 61) ^ (a >> 16)
	a = a + (a << 3)
	a = a ^ (a >> 4)
	a = a * 0x27d4eb2d
	a = a ^ (a >> 15)
	return a
}

func CheckAccount(userName string, passwd string) (int64, string) {

	aUrl := fmt.Sprintf("http://127.0.0.1:9090/v1/accounts?user_name=%s&passwd=%s",
		userName, passwd)

	aRsp := &CheckAccoutRsp{}
	cost, err := ttcomm.GetApi(aUrl, "", aRsp)
	if err != nil {
		panic(err)
	}

	//fmt.Println(aRsp)

	return cost, aRsp.Token

}

var userIDs [20]uint32
var teamIDs [2]uint32
var articleIDs [400]uint32
var commentIDs [1000]uint32
var scoreIDs [1000]uint32

func tableIDAutoIDToID(tableID, autoID uint32) (ID uint32) {
	ID = autoID
	ID <<= 10
	ID += tableID
	return ID
}

func MakeUserID(userName string) uint32 {
	tableID := uint32(BKDRHash(userName) % 20)
	userIDs[tableID]++
	return tableIDAutoIDToID(tableID, userIDs[tableID])
}

func MakeArticleID(userID uint32) uint32 {
	tableID := (hash32(userID) % 400)
	articleIDs[tableID]++
	return tableIDAutoIDToID(tableID, articleIDs[tableID])
}

func GetTop(token string) int64 {
	url := fmt.Sprintf("http://127.0.0.1:9090/v1/top")
	fmt.Println(url)

	rsp := &TopListWrap{}
	cost, err := ttcomm.GetApi(url, token, rsp)
	if err != nil {
		panic(err)
	}

	fmt.Println(rsp)

	return cost

}

func GetTeam(teamName string, token string) int64 {
	tableID := uint32(BKDRHash(teamName) % 2)
	teamIDs[tableID]++
	teamID := tableIDAutoIDToID(tableID, teamIDs[tableID])

	url := fmt.Sprintf("http://127.0.0.1:9090/v1/teams/%d", teamID)
	//	fmt.Println(url)

	rsp := &Team{}
	cost, err := ttcomm.GetApi(url, token, rsp)
	if err != nil {
		panic(err)
	}

	return cost

}

func GetArticle(articleID uint32, token string) int64 {

	url := fmt.Sprintf("http://127.0.0.1:9090/v1/articles/%d", articleID)

	rsp := &Article{}
	cost, err := ttcomm.GetApi(url, token, rsp)
	if err != nil {
		panic(err)
	}

	//	fmt.Println("article: ", rsp)

	return cost

}

func GetComment(articleID uint32, token string) int64 {

	url := fmt.Sprintf("http://127.0.0.1:9090/v1/articles/%d/comments", articleID)

	var rsps []Comment
	cost, err := ttcomm.GetApi(url, token, &rsps)
	if err != nil {
		//		fmt.Println(err)
	}

	//	fmt.Println("comment:", rsps)

	return cost

}

func GetScore(articleID uint32, token string) int64 {

	url := fmt.Sprintf("http://127.0.0.1:9090/v1/articles/%d/score", articleID)

	rsp := &Score{}
	cost, err := ttcomm.GetApi(url, token, rsp)
	if err != nil {
		panic(err)
	}

	//fmt.Println("score: ", rsp)

	return cost

}

func CreateAccount(name string, passwd string) (int64, uint32) {

	req := CreateAccountReq{Name: name, Passwd: passwd}
	var rsp = &IDResp{}
	cost, err := ttcomm.PostApi("http://127.0.0.1:9090/v1/accounts", "", req, rsp)
	if err != nil {
		panic(err)
	}

	return cost, uint32(rsp.ID)

}

func CreateTeam(name string, desc string, userID uint32, token string) (int64, uint32) {
	req := CreateTeamReq{Name: name, Desc: name, UserID: userID}
	var rsp = &IDResp{}
	cost, err := ttcomm.PostApi("http://127.0.0.1:9090/v1/teams", token, req, rsp)
	if err != nil {
		panic(err)
		return 0, 0
	}

	//fmt.Println(rsp)

	return cost, uint32(rsp.ID)
}

func CreateArticle(name string, userID uint32, content string, token string) (int64, uint32) {
	req := CreateArticleReq{Title: name, UserID: userID, Content: content}
	var rsp = &IDResp{}
	cost, err := ttcomm.PostApi("http://127.0.0.1:9090/v1/articles", token, req, rsp)
	if err != nil {
		panic(err)
		return 0, 0

	}

	return cost, uint32(rsp.ID)
}

func CreateComment(articleID uint32, userID uint32, comment string, token string) (int64, uint64) {
	req := CreateCommentReq{ArticleID: articleID, UserID: userID, Content: comment}
	url := fmt.Sprintf("http://127.0.0.1:9090/v1/articles/%d/comments", articleID)
	var rsp = &IDResp{}
	cost, err := ttcomm.PostApi(url, token, req, rsp)
	if err != nil {
		panic(err)
		return 0, 0
	}

	return cost, rsp.ID

}

func UpdateScore(articleID uint32, userID uint32, score uint32, token string) int64 {
	req := UpdateScoreReq{ArticleID: articleID, UserID: userID, Score: score}
	url := fmt.Sprintf("http://127.0.0.1:9090/v1/articles/%d/score", articleID)
	var rsp = &ttcomm.CommResp{}
	cost, err := ttcomm.PostApi(url, token, req, rsp)
	if err != nil {
		panic(err)
		return 0
	}

	return cost
}
