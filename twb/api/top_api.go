package tapi

import (
	"net/http"
	"tlogic"
	"tutil"
)

func GetTop(w http.ResponseWriter, r *http.Request) {
	tutil.LogInfo("GetTop")

	err := checkToken(r)
	if err != nil {
		httpBadRequest(w)
		return
	}

	tpList, err := tlogic.GetTop()
	if err == tutil.ErrNotFound {
		httpBadRequest(w)
		return
	}

	if err == tutil.InternalErr {
		httpInterErr(w)
		return
	}

	httpResp(w, tpList)

}
