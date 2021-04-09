package tdao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"math/rand"
	"time"
	"tutil"
)

var db *sqlx.DB

func hash32(a uint32) (b uint32) {
	a = (a ^ 61) ^ (a >> 16)
	a = a + (a << 3)
	a = a ^ (a >> 4)
	a = a * 0x27d4eb2d
	a = a ^ (a >> 15)
	return a
}

func iDToTableIDAutoID(ID uint32) (tableID, autoID uint32) {
	tableID = ID & 0x3FF
	autoID = (ID >> 10)
	return tableID, autoID
}

func tableIDAutoIDToID(tableID, autoID uint32) (ID uint32) {
	ID = autoID
	ID <<= 10
	ID += tableID
	return ID
}

func BKDRHash(str string) (hash uint32) {
	seed := uint32(131) // 31 131 1313 13131 131313 etc..
	hash = uint32(0)
	for i := 0; i < len(str); i++ {
		hash = (hash * seed) + uint32(str[i])
	}
	return hash
}

func GetRow(dest interface{}, query string, args ...interface{}) (err error) {
	err = db.Get(dest, query, args...)
	if err == sql.ErrNoRows {
		tutil.Info.Println("NoRows CheckRow query %s", query)
		return tutil.ErrNotFound
	}
	if err != nil {
		tutil.Err.Println("CheckRow query %s", query, args, err)
		return tutil.InternalErr
	}

	return nil

}

func GetRows(dest interface{}, query string, args ...interface{}) (err error) {
	err = db.Select(dest, query, args...)
	if err == sql.ErrNoRows {
		tutil.Info.Println("NoRows GetRows query %s", query)
		return tutil.ErrNotFound
	}
	if err != nil {
		tutil.Err.Println("GetRows query %s", query)
		return tutil.InternalErr
	}

	return nil
}

func InsertRow32ID(tableID uint32, query string, args ...interface{}) (uint32, error) {

	res, err := db.Exec(query, args...)
	if err != nil {
		tutil.Err.Println("InsertRow", query, err)
		return 0, err
	}

	var autoID int64
	autoID, err = res.LastInsertId()
	if err != nil {
		tutil.Err.Println("InsertRow", query, err)
		return 0, err
	}

	tutil.Info.Println("InsertRow32ID", tableID, autoID, tableIDAutoIDToID(tableID, uint32(autoID)))

	return tableIDAutoIDToID(tableID, uint32(autoID)), err
}

func InsertRow64ID(tableID uint32, query string, args ...interface{}) (uint64, error) {

	res, err := db.Exec(query, args...)
	if err != nil {
		tutil.Err.Println("InsertRow", query, err)
		return 0, err
	}

	var ID int64
	ID, err = res.LastInsertId()
	if err != nil {
		tutil.Err.Println("InsertRow", query, err)
		return 0, err
	}

	return uint64(ID), err
}

func UpdateRow(tableID uint32, query string, args ...interface{}) error {
	res, err := db.Exec(query, args...)
	if err != nil {
		tutil.Err.Println("InsertRow", query, err)
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		tutil.Err.Println("UpdateRow", query, err)
		return err
	}

	return nil

}

func init() {

	dsn := "entrytask:Twb@123456@/twb_db"

	var err error
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		panic(err)
	}

	rand.Seed(time.Now().UnixNano())
	fmt.Println(db)

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)

}
