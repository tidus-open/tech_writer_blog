package tdao

import (
	//"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"tutil"
)

type Account struct {
	ID      uint32 `db:"idx_user_id"`
	Name    string `db:"user_name"`
	Passwd  string `db:"passwd"`
	Delflag uint32 `db:"delflag"`
}

func CheckAccount(name string, passwd string) (err error) {
	tableID := uint32(BKDRHash(name) % 20)

	tutil.Info.Println("CheckAccount name %s tableID %v passwd %v", name, tableID, passwd)

	query := fmt.Sprintf("select idx_user_id, user_name, passwd, delflag from twb_account_tab_%08d where user_name = ? and passwd = ?", tableID)
	var acct Account
	return GetRow(&acct, query, name, passwd)

}

func CreateAccount(name string, passwd string) (uint32, error) {
	tableID := uint32(BKDRHash(name) % 20)

	fmt.Printf("CreateAccount tableID %v \n", tableID)

	query := fmt.Sprintf("insert into twb_account_tab_%08d(user_name, passwd, delflag) values(?,?,?)", tableID)

	return InsertRow32ID(tableID, query, name, passwd, 0)

}
