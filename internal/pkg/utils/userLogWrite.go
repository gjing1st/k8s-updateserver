package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	HttpCtJson = "application/json"
)

type Reader []byte

func (r *Reader) Read(p []byte) (n int, err error) {
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

func UserLogWrite(req interface{}, d time.Duration) (sc int, err error) {
	var reader io.Reader
	var resp *http.Response
	var data Reader
	sc = -1
	if req != nil {
		data, err = json.Marshal(req)
		fmt.Printf("reqData==============%#v\n", string(data))

		if err != nil {
			return
		}
		reader = &data
	}
	client := &http.Client{
		Timeout: d,
	}

	// 测试的时候需要注释掉下面的一条
	// fmt.Printf("Inner Url :: %s \n",K8sConfig.Admin.Url)
	url := K8sConfig.Backend.Addr + "/tna-cipher/v1/log"

	resp, err = client.Post(url, HttpCtJson, reader)

	if err != nil {
		fmt.Printf("req faild :: %s \n", err)
		return
	}
	defer resp.Body.Close()
	sc = resp.StatusCode
	fmt.Printf("return sc :: %d \n", sc)
	return
}
