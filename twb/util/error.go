package tutil

import (
	"errors"
)

//api
type ErrType int32

const (
	OK                ErrType = 0
	CommAuthFail      ErrType = -1
	AcctInValidMail   ErrType = -2
	AcctInvalidPasswd ErrType = -3
)

//logic
var ErrNotFound = errors.New("data not found")
var InternalErr = errors.New("InternalErr")

//dao
