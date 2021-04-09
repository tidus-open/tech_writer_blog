package tdao

import (
	"fmt"
	"testing"
)

/*
func TestCreateTeam(t *testing.T) {
	db, err := initDB()
	if err != nil {
		t.Errorf(`%q`, err)
		return
	}

	team := Team{}
	team.db = db
	team.TeamID = 1
	var id uint32
	id, err = team.CreateTeam("test", "test", 1)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(id)
	fmt.Println(iDToTableIDAutoID(id))

	t.Error(`test....`)
}
*/

func TestCheckAccount(t *testing.T) {
	err := CheckAccount("delphi", "123456")
	fmt.Println(err)
}
