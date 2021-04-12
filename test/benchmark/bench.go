package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
	"ttcomm"
	"ttunit"
)

type Stat struct {
	MaxCost   int64 //ns
	TotalCost int64 //ns
	Lt100ms   int64
	TotalCnt  int64
}

func max(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

func (s *Stat) DoStat(cost int64) {

	s.TotalCnt++
	s.MaxCost = max(s.MaxCost, cost)
	s.TotalCost += cost
	if cost < int64(1e8) {
		s.Lt100ms++
	}

}

func Write(st *Stat, userName string, passwd string) {

	fmt.Println("write ", userName)

	cost, userID := ttunit.CreateAccount(userName, passwd)
	st.DoStat(cost)

	fmt.Println("w userID :", userID)

	for i := 0; i <= 14; i++ {

		articleName := userName + "article1" + strconv.Itoa(i)

		var articleID uint32
		cost, articleID = ttunit.CreateArticle(articleName, userID, articleName, "")
		st.DoStat(cost)

		//		fmt.Println("w articleID : ", articleID)

		cost, _ = ttunit.CreateComment(articleID, userID, articleName, "")
		st.DoStat(cost)

		cost = ttunit.UpdateScore(articleID, userID, 3, "")
		st.DoStat(cost)

		teamName := userName + "team" + strconv.Itoa(i)

		cost, _ = ttunit.CreateTeam(teamName, teamName, userID, "")
		st.DoStat(cost)
	}

}

func WriteRange(st *Stat, start int, end int, done func()) {
	defer done()
	for i := start; i < end; i++ {
		idx := strconv.Itoa(i)
		Write(st, idx, idx)
	}
}

func StressTestWrite(totalCnt int, procs int) {
	var st Stat
	var wg sync.WaitGroup

	cntPerProc := totalCnt / procs

	sTime := (time.Now().UnixNano() / 1e9)
	wg.Add(procs)
	for i := 0; i < procs; i++ {
		start := i * cntPerProc
		end := start + cntPerProc
		go WriteRange(&st, start, end, wg.Done)
	}
	wg.Wait()

	eTime := (time.Now().UnixNano() / 1e9)

	ttcomm.Info.Println("total Time:", eTime-sTime, "s", "| count: ", totalCnt+totalCnt*15*4, "| max resp time: ", st.MaxCost/1e6, "ms", "| total cost : ", st.TotalCost/1e6, "ms ", "| avg resp time: ", st.TotalCost/1e6/st.TotalCnt, "ms", "| P99: ", float32(st.Lt100ms)/float32(totalCnt+totalCnt*15*4), " | < 100ms : ", st.Lt100ms)

}

func Read(st *Stat, userName string, passwd string) {
	fmt.Println("read: ", userName)

	cost, _ := ttunit.CheckAccount(userName, passwd)
	st.DoStat(cost)

	userID := ttunit.MakeUserID(userName)
	//	fmt.Println("r userID :", userID)

	for i := 0; i <= 14; i++ {

		teamName := userName + "team" + strconv.Itoa(i)

		cost = ttunit.GetTeam(teamName, "")
		st.DoStat(cost)

		//articleName := userName + "article1" + strconv.Itoa(i)

		articleID := ttunit.MakeArticleID(userID)
		//		fmt.Println("r articleID :", articleID)

		cost = ttunit.GetArticle(articleID, "")
		st.DoStat(cost)

		cost = ttunit.GetComment(articleID, "")
		st.DoStat(cost)

		cost = ttunit.GetScore(articleID, "")
		st.DoStat(cost)

	}

}

func ReadRange(st *Stat, start int, end int, done func()) {
	defer done()
	for i := start; i < end; i++ {
		idx := strconv.Itoa(i)
		Read(st, idx, idx)
	}
}

func StressTestRead(totalCnt int, procs int) {
	var st Stat
	var wg sync.WaitGroup

	cntPerProc := totalCnt / procs

	sTime := (time.Now().UnixNano() / 1e9)
	wg.Add(procs)
	for i := 0; i < procs; i++ {
		start := i * cntPerProc
		end := start + cntPerProc
		go ReadRange(&st, start, end, wg.Done)
	}
	wg.Wait()

	eTime := (time.Now().UnixNano() / 1e9)

	ttcomm.Info.Println("total Time:", eTime-sTime, "s", "| count: ", totalCnt+totalCnt*15*4, "| max resp time: ", st.MaxCost/1e6, "ms", "| total cost : ", st.TotalCost/1e6, "ms ", "| avg resp time: ", st.TotalCost/1e6/st.TotalCnt, "ms", "| P99: ", float32(st.Lt100ms)/float32(totalCnt+totalCnt*15*4), " | < 100ms : ", st.Lt100ms)

}

func StreeWriteAndRead(totalCnt int, procs int) {
	var st Stat
	var wg sync.WaitGroup

	cntPerProc := totalCnt / procs

	sTime := (time.Now().UnixNano() / 1e9)
	wg.Add(procs)

	fmt.Println("write start")
	for i := 0; i < int(procs/2); i++ {
		start := i * cntPerProc
		end := start + cntPerProc
		go WriteRange(&st, start, end, wg.Done)
	}

	time.Sleep(10 * time.Second)

	fmt.Println("read start")

	for i := 0; i < int(procs/2)+int(procs%2); i++ {
		start := i * cntPerProc
		end := start + cntPerProc
		go ReadRange(&st, start, end, wg.Done)
	}
	wg.Wait()

	eTime := (time.Now().UnixNano() / 1e9)

	ttcomm.Info.Println("total Time:", eTime-sTime, "s", "| count: ", totalCnt+totalCnt*15*4, "| max resp time: ", st.MaxCost/1e6, "ms", "| total cost : ", st.TotalCost/1e6, "ms ", "| avg resp time: ", st.TotalCost/1e6/st.TotalCnt, "ms", "| P99: ", float32(st.Lt100ms)/float32(totalCnt+totalCnt*15*4), " | < 100ms : ", st.Lt100ms)

}
