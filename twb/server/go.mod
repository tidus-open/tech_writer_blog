module server

go 1.16

require (
	github.com/gorilla/mux v1.8.0
	tapi v0.0.0-00010101000000-000000000000 // indirect
	tdao v0.0.0-00010101000000-000000000000 // indirect
	tlogic v0.0.0-00010101000000-000000000000 // indirect
	tutil v0.0.0-00010101000000-000000000000 // indirect
)

replace tapi => ../api

replace tlogic => ../logic

replace tdao => ../dao

replace tutil => ../util
