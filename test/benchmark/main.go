package main

func main() {

	//	CreateAccount("1", "1")
	//	CheckAccount("1", "1")

	//	CreateTeam("team1", "team1", 12345)
	//	GetTeam("team1")

	//	CreateArticle("article1", 123456, "article1")
	//	GetArticle("article1", 123456)

	//	CreateComment(12345, 12345, "comment1")
	//	GetComment(12345)

	//	UpdateScore(12345, 12345, 3)
	//	GetScore(12345)

	//	var st Stat
	//	Write(&st, "bbb", "bbb")
	//	Read(&st, "bbb", "bbb")

	//	tableID := uint32(hash32(1025) % 400)
	//	fmt.Println(tableID)
	//	articleIDs[tableID]++
	//	articleID := tableIDAutoIDToID(tableID, articleIDs[tableID])
	//	fmt.Println(articleID)

	//	articleID := makeArticleID(1025)
	//	fmt.Println(articleID)

	//	StressTestWrite(100000, 45)
	//	StressTestRead(100000, 30)
	//	fmt.Println(uint32(1e6))
	StreeWriteAndRead(100000, 45)

	//	cost, userID := CreateAccount("delpdi", "delphi")
	//	fmt.Println(cost/1e6, " ", userID)

	//GetTop("")
	//	go WriteForTop()
	//	go ShowTop()
	//
	//	time.Sleep(10000 * time.Second)

	//	CreateAccount("delphi", "delphi")
	//_, token := CheckAccount("delphi", "delphi")
	//CreateTeam("delphi", "delphi", 12345, token)

}
