package tutil

import (
	"io"
	"log"
	"os"
)

var (
	Info *log.Logger
	Warn *log.Logger
	Err  *log.Logger
)

func init() {
	infoFile, err := os.OpenFile("../../log/info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	warnFile, err := os.OpenFile("../../log//warning.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	errFile, err := os.OpenFile("../../log/error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	Info = log.New(io.MultiWriter(os.Stdout, infoFile), "Info:", log.Ldate|log.Ltime|log.Lshortfile)
	Warn = log.New(io.MultiWriter(os.Stdout, warnFile), "Warning:", log.Ldate|log.Ltime|log.Lshortfile)
	Err = log.New(io.MultiWriter(os.Stdout, errFile), "Error:", log.Ldate|log.Ltime|log.Lshortfile)

}
