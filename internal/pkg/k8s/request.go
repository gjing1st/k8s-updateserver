package k8s

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"syscall"
	"time"
)

var (
	HttpCtJson      = "application/json" //json格式
	HttpXUrlencoded = "application/x-www-form-urlencoded"
)

// TokenRequestTimeout
// @description: 请求k8s获取token
// @param:
// @author: GJing
// @email: guojing@tna.cn
// @date: 2022/7/27 17:12
// @success:
func TokenRequestTimeout(method, url string, reqData url.Values, res interface{}, d time.Duration) (sc int, err error) {
	var resp *http.Response
	sc = -1

	client := &http.Client{
		Timeout: d,
	}
	if method == "GET" {
		resp, err = client.Get(url)
	} else if method == "POST" {
		resp, err = client.Post(url, HttpXUrlencoded, strings.NewReader(reqData.Encode()))
	} else {
		var req *http.Request
		req, err = http.NewRequest(method, url, strings.NewReader(reqData.Encode()))
		if err != nil {
			return
		}
		req.Header.Set("Content-Type", HttpXUrlencoded)
		resp, err = client.Do(req)
	}
	if err != nil {
		return
	}
	defer resp.Body.Close()
	sc = resp.StatusCode
	if (sc != http.StatusNotAcceptable) && (sc < http.StatusOK || sc >= http.StatusMultipleChoices) {
		err = syscall.EINVAL
		return
	}
	if res != nil {
		data, err1 := ioutil.ReadAll(resp.Body)
		if err1 != nil {
			return sc, err1
		}
		err = json.Unmarshal(data, res)
		if err != nil {
			fmt.Println("json.Unmarshal err", err)
		}
	}
	return
}

type ReadCloser []byte

func (r *ReadCloser) Read(p []byte) (n int, err error) {
	t := len(p)
	l := len(*r)
	if t == 0 {
		return
	} else if l == 0 {
		err = io.EOF
	} else {
		copy(p, *r)
		if t < l {
			n = t
		} else {
			n = l
		}
		*r = (*r)[n:l]
	}
	return
}
func (r *ReadCloser) Close() (err error) {
	return
}

func JsonRequestTimeout(method,url string, reqData url.Values, res interface{}, d time.Duration) (sc int, err error) {
	token, err1 := GetToken()
	if err1 != nil {
		return 0, err
	}
	//var reader io.ReadCloser
	var resp *http.Response
	var data ReadCloser
	sc = -1
	//if reqData != nil {
	//	data, err = json.Marshal(reqData)
	//	if err != nil {
	//		return
	//	}
	//}
	client := &http.Client{
		Timeout: d,
	}
	var req *http.Request
	req, err = http.NewRequest(method, url, strings.NewReader(reqData.Encode()))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", HttpCtJson)
	req.Header.Set("Authorization", token)
	resp, err = client.Do(req)
	if err != nil {
		fmt.Println("err======",err)
		return
	}
	defer resp.Body.Close()
	sc = resp.StatusCode
	if sc != http.StatusOK{
		log.WithFields(log.Fields{"http code ": sc}).Error("请求接口返回状态不是200")
		return
	}


	if res != nil {
		data, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return
		}
		err = json.Unmarshal(data, res)
		if err != nil {
			fmt.Println("json.Unmarshal err", err)
		}
	}
	return
}



// JsonRestRequestTimeout json格式请求
func JsonRestRequestTimeout(method, url string, req, res interface{}, d time.Duration) (sc int, err error) {
	token, err1 := GetToken()
	if err1 != nil {
		return 0, err
	}
	var reader io.ReadCloser
	var resp *http.Response
	var data ReadCloser
	sc = -1
	if req != nil {
		data, err = json.Marshal(req)
		fmt.Println("-------",string(data))
		if err != nil {
			return
		}
		reader = &data
	}
	client := &http.Client{
		Timeout: d,
	}
	if method == "GET" {
		resp, err = client.Get(url)
	}  else {
		var req *http.Request
		fmt.Println("url",url)
		fmt.Println("method",method)
		req, err = http.NewRequest(method, url, reader)
		if err != nil {
			return
		}
		req.Header.Set("Content-Type", HttpCtJson)
		req.Header.Set("Authorization", token)
		resp, err = client.Do(req)
	}
	if err != nil {
		return
	}
	defer resp.Body.Close()
	sc = resp.StatusCode
	if (sc != http.StatusNotAcceptable) && (sc < http.StatusOK || sc >= http.StatusMultipleChoices) {
		err = syscall.EINVAL
		return
	}
	if res != nil {
		data, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return
		}
		err = json.Unmarshal(data, res)
		if err != nil {
			fmt.Println("json.Unmarshal err", err)
		}
	}
	return
}