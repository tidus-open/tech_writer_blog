package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"testing"
)

func CheckResp(t *testing.T, resp *http.Response, wantStatusCode int, wantBody interface{}) {
	if resp.StatusCode != wantStatusCode {
		t.Errorf("statusCode %v != %v ", resp.StatusCode, wantStatusCode)
		//panic("")
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
		t.Errorf("body %v != %v", rsp, wantBody)
	}

}

func makeGetReq(url string) *http.Request {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	//defer request.Body.Close()
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	request.Header.Set("Token", "123456789")

	return request

}

func CheckGetApi(t *testing.T, url string, wantStatusCode int, wantBody interface{}) {
	req := makeGetReq(url)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	CheckResp(t, resp, wantStatusCode, wantBody)
}

func makePostReq(url string, src interface{}) *http.Request {
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

	return request

}

func CheckPostApi(t *testing.T, url string, src interface{}, wantStatusCode int, wantBody interface{}) {
	req := makePostReq(url, src)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	CheckResp(t, resp, wantStatusCode, wantBody)

}

func PostApi(t *testing.T, url string, src interface{}, dest interface{}) error {
	req := makePostReq(url, src)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("1", err)
		return err
	}
	defer resp.Body.Close()

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
		fmt.Println("4", err)
		return err
	}

	return nil

}

var client = &http.Client{}

func init() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client = &http.Client{
		Transport: tr,
	}

}
