package ttcomm

import (
	"bytes"
	//"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type CommResp struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

var Info *log.Logger

func CheckResp(resp *http.Response, wantStatusCode int, wantBody interface{}) {
	if resp.StatusCode != wantStatusCode {
		fmt.Println("statusCode %v != %v ", resp.StatusCode, wantStatusCode)
	}

	body, err := ioutil.ReadAll(io.LimitReader(resp.Body, 10240))
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := resp.Body.Close(); err != nil {
		fmt.Println(err)
		return
	}

	var rsp CommResp
	if err := json.Unmarshal(body, &rsp); err != nil {
		fmt.Println(err)
		return
	}

	if wantBody != nil && rsp != wantBody {
		fmt.Println("body %v != %v", rsp, wantBody)
	}

}

func makeGetReq(url string, token string) *http.Request {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	//defer request.Body.Close()
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	request.Header.Set("Token", token)

	return request

}

func CheckGetApi(url string, token string, wantStatusCode int, wantBody interface{}) {
	req := makeGetReq(url, token)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	CheckResp(resp, wantStatusCode, wantBody)
}

func makePostReq(url string, token string, src interface{}) *http.Request {
	bytesData, err := json.Marshal(&src)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	reader := bytes.NewReader(bytesData)

	request, err := http.NewRequest("POST", url, reader)
	defer request.Body.Close()
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	request.Header.Set("Token", token)
	//	fmt.Println("token:", token)

	return request

}

func CheckPostApi(url string, token string, src interface{}, wantStatusCode int, wantBody interface{}) {
	req := makePostReq(url, token, src)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	CheckResp(resp, wantStatusCode, wantBody)

}

func GetResp(resp *http.Response, dest interface{}) error {
	body, err := ioutil.ReadAll(io.LimitReader(resp.Body, 10240))
	if err != nil {
		fmt.Println("2", err)
		return err
	}

	if err := resp.Body.Close(); err != nil {
		fmt.Println("3", err)
		return err
	}

	if err := json.Unmarshal(body, dest); err != nil {
		//		fmt.Println("4", err)
		return err
	}

	return nil
}

func PostApi(url string, token string, src interface{}, dest interface{}) (int64, error) {
	req := makePostReq(url, token, src)
	sTime := (time.Now().UnixNano())
	resp, err := client.Do(req)
	eTime := (time.Now().UnixNano())
	if err != nil {
		fmt.Println("1", err)
		return 0, err
	}
	defer resp.Body.Close()

	err = GetResp(resp, dest)
	return eTime - sTime, err
}

func GetApi(url string, token string, dest interface{}) (int64, error) {
	req := makeGetReq(url, token)
	sTime := (time.Now().UnixNano())
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	eTime := (time.Now().UnixNano())
	err = GetResp(resp, dest)
	if err != nil {
		//	fmt.Println(err)
		return 0, err
	}
	defer resp.Body.Close()

	return eTime - sTime, err

}

var client *http.Client

func init() {
	//	tr := &http.Transport{
	//		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	//	}
	//	client = &http.Client{
	//		Transport: tr,
	//	}

	infoFile, err := os.OpenFile("info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	Info = log.New(infoFile, "Info:", log.Ldate|log.Ltime|log.Lshortfile)

	defaultRoundTripper := http.DefaultTransport
	defaultTransportPointer, ok := defaultRoundTripper.(*http.Transport)
	if !ok {
		panic(fmt.Sprintf("defaultRoundTripper not an *http.Transport"))
	}
	defaultTransport := *defaultTransportPointer // dereference it to get a copy of the struct that the pointer points to
	defaultTransport.MaxIdleConns = 100
	defaultTransport.MaxIdleConnsPerHost = 100

	client = &http.Client{Transport: &defaultTransport}

}
