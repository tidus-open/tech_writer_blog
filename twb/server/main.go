package main

import (
	//	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	//	"tdao"
	//	"tlogic"
)

func main() {

	/*
		tdao.TAddItem(1, 111)
		tdao.TAddItem(2, 222)
		tdao.TAddItem(3, 333)

		tlogic.UpdateTopList(4, 444)
		tlogic.UpdateTopList(5, 555)
		tlogic.UpdateTopList(6, 666)
		tlogic.UpdateTopList(7, 777)
		tlogic.UpdateTopList(8, 888)
		tlogic.UpdateTopList(9, 999)
		tlogic.UpdateTopList(10, 100)
		tlogic.UpdateTopList(11, 111)
		tlogic.UpdateTopList(12, 222)

		tpList, err := tlogic.GetTop()
		fmt.Println(tpList, err)
	*/

	router := NewRouter()
	//	go tlogic.RefreshTopList()

	router.PathPrefix("/debug/pprof/").Handler(http.DefaultServeMux)
	//	log.Fatal(http.ListenAndServeTLS(":9090", "key/cert.pem", "key/key.pem", router))
	log.Fatal(http.ListenAndServe("0.0.0.0:9090", router))

}
