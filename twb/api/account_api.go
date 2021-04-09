package tapi

import (
	"fmt"
	"net/http"
	"tlogic"
	"tutil"
)

type CreateAccountRsp struct {
	ID uint32 `json:"user_id"`
}

type Account struct {
	Name   string `json:"user_name"`
	Passwd string `json:"passwd"`
}

//for login
type CheckAccountResp struct {
	Token string `json:"token"`
}

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	var acct = &tlogic.Account{}

	err := checkToken(r)
	if err != nil {
		httpBadRequest(w)
		return
	}

	err = GetJsonParam(acct, r)
	if err != nil {
		httpBadRequest(w)
		return
	}

	fmt.Println("CreateAcount", acct.Name, acct.Passwd)

	acct.Passwd = getMd5String(getMd5String(acct.Passwd+"dxetewg") + "vebet")

	var autoID uint32
	if autoID, err = tlogic.CreateAccount(acct.Name, acct.Passwd); err != nil {
		httpBadRequest(w)
		return
	}

	rsp := CreateAccountRsp{ID: autoID}
	httpResp(w, rsp)

}

func CheckAccount(w http.ResponseWriter, r *http.Request) {
	tutil.Info.Println("checkAccount")

	name, _ := getParamStr(r.URL.Query(), "user_name")
	passwd, _ := getParamStr(r.URL.Query(), "passwd")

	err := checkToken(r)
	if err != nil {
		httpBadRequest(w)
		return
	}

	passwd = getMd5String(getMd5String(passwd+"dxetewg") + "vebet")

	err = tlogic.CheckAccount(name, passwd)
	if err == tutil.ErrNotFound {
		httpBadRequest(w)
		return
	}

	if err == tutil.InternalErr {
		httpInterErr(w)
		return
	}

	token := generateToken(name, passwd)

	fmt.Println(token)

	rsp := CheckAccountResp{Token: token}
	httpResp(w, rsp)

}
