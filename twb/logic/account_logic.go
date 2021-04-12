package tlogic

import (
	"tdao"
	"tutil"
)

type Account struct {
	Name   string `json:"user_name"`
	Passwd string `json:"passwd"`
}

func CreateAccount(name string, passwd string) (uint32, error) {
	tutil.LogInfo(" CreateAccount", name, passwd)
	return tdao.CreateAccount(name, passwd)
}

func CheckAccount(name string, passwd string) (err error) {
	tutil.LogInfo("CheckAccount", name, passwd)
	return tdao.CheckAccount(name, passwd)
}
