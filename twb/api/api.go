package tapi

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"tutil"
)

type Resp struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

type autoIDResp struct {
	id uint32 `json:"id"`
}

func httpBadRequest(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusBadRequest) // unprocessable entity
	rsp := Resp{Code: "INVALID_PARAM", Msg: "invalid param"}
	tutil.Err.Println("httpBadRequest", rsp)
	if err := json.NewEncoder(w).Encode(rsp); err != nil {
		panic(err)
	}

}

func httpInterErr(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusInternalServerError)
	rsp := Resp{Code: "INTERNAL_ERROR", Msg: "internal server error"}
	if err := json.NewEncoder(w).Encode(rsp); err != nil {
		panic(err)
	}

}

func httpNotFound(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	rsp := Resp{Code: "NOT_FOUND", Msg: "not found"}
	if err := json.NewEncoder(w).Encode(rsp); err != nil {
		panic(err)
	}

}

func httpOK(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	rsp := Resp{Code: "OK", Msg: "ok"}
	if err := json.NewEncoder(w).Encode(rsp); err != nil {
		panic(err)
	}

}

func httpResp(w http.ResponseWriter, rsp interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&rsp); err != nil {
		panic(err)
	}

}

func getParamStr(vars url.Values, name string) (value string, isExist bool) {
	param, ok := vars[name]
	if !ok {
		return "", false
	} else {
		fmt.Printf("%s : %s\n", name, param[0])
		return param[0], true
	}
}

func getParamUInt(vars url.Values, name string) (uint32, bool) {
	strval, ok := getParamStr(vars, name)
	if !ok {
		return 0, false
	}

	intval, err := strconv.Atoi(strval)
	if err != nil {
		return 0, false
	}

	return uint32(intval), true
}

func checkToken(r *http.Request) (err error) {
	fmt.Println("token :", r.Header["Token"])
	return nil
}

func GetJsonParam(dest interface{}, r *http.Request) error {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1024))
	if err != nil {
		tutil.Err.Println(err)
		return err
	}
	fmt.Println(body)
	if err := r.Body.Close(); err != nil {
		tutil.Err.Println(err)
		return err
	}
	if err := json.Unmarshal(body, dest); err != nil {
		tutil.Err.Println(err)
		return err
	}

	return nil

}

func ProcessPost(w http.ResponseWriter, r *http.Request, dest interface{}) {

}
