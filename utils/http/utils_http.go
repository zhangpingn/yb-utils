// Copyright (c) 2023 YaoBase
// 描述: 封装http请求相关方法
// 创建者: 张平
// 创建时间: 2022/5/26 17:13

package http

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/zyzp5217758/YBUtils/log"
	"github.com/zyzp5217758/YBUtils/model/result"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"strings"
)

// Get
// 描述：封装get请求
// 创建者：张平
// 创建时间：2022/5/26 17:14
// 参数：
//	 [in]
//		requestUrl：请求的url
//		token：请求携带的token信息
//   [out]
//		body：请求资源返回的数据
//		err：请求过程中出现的错误
//   [in/out]
//		无
func Get(requestUrl, token string) (body []byte, err error) {
	// 跳过证书验证
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	// http cookie接口
	cookieJar, _ := cookiejar.New(nil)
	c := &http.Client{
		Jar:       cookieJar,
		Transport: tr,
	}
	var response *http.Response

	req, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		return nil, err
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("Accept", "application/json;charset=UTF-8")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("X-Auth-Token", token)

	// 重连2次
	for i := 0; i < 2; i++ {
		response, err = c.Do(req)
		if err != nil {
			times := i + 1
			log.Error(fmt.Sprintf("connect times: (%d) net connect has error: %s", times, err.Error()), zap.String("requestUrl", requestUrl))
		} else {
			break
		}
	}
	if response != nil {
		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, errors.New("api request has error: " + err.Error())
		}
		return body, nil
	}
	if err != nil {
		return nil, errors.New("api request has error: " + err.Error())
	}
	return nil, errors.New("api request has error: err = nil")
}

// doPost
// 描述：封装的请求处理方法
// 创建者：张平
// 创建时间：2022/5/26 17:14
// 参数：
//	 [in]
//		action：请求的方法
//		requestUrl：请求的url
//		token：请求携带的token信息
//		param：请求参数
//   [out]
//		authToken：认证token
//		body：请求返回的数据
//		err：请求过程中出现的错误
//   [in/out]
//		无
func doPost(action, requestUrl, token, param string) (authToken string, res *result.Result, err error) {
	// 跳过证书验证
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	// http cookie接口
	cookieJar, _ := cookiejar.New(nil)
	c := &http.Client{
		Jar:       cookieJar,
		Transport: tr,
	}

	var response *http.Response

	req, err := http.NewRequest(action, requestUrl, strings.NewReader(param))
	if err != nil {
		return "", nil, err
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("Accept", "application/json;charset=UTF-8")
	req.Header.Set("Connection", "keep-alive")
	if token != "" {
		req.Header.Set("X-Auth-Token", token)
	}

	// 重连2次
	for i := 0; i < 2; i++ {
		response, err = c.Do(req)
		if err != nil {
			times := i + 1
			log.Error(fmt.Sprintf("connect times: (%d) net connect has error: %s", times, err.Error()), zap.String("requestUrl", requestUrl), zap.String("param", param))
		} else {
			break
		}
	}
	if response != nil {
		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return "", nil, errors.New("api request has error: " + err.Error())
		}
		res = &result.Result{}
		if err = json.Unmarshal(body, &res); err != nil {
			return response.Header.Get("X-Subject-Token"), nil, err
		}

		return response.Header.Get("X-Subject-Token"), res, nil
	}
	if err != nil {
		return "", nil, errors.New("api request has error: " + err.Error())
	} else {
		return "", nil, errors.New("api request has error: err = nil")
	}
}

// Post
// 描述：封装https的post方法
// 创建者：张平
// 创建时间：2022/5/26 17:14
// 参数：
//	 [in]
//		requestUrl：请求的url
//		token：请求携带的token信息
//		param：请求参数
//   [out]
//		authToken：认证token
//		body：请求返回的数据
//		err：请求过程中出现的错误
//   [in/out]
//		无
func Post(requestUrl, token, param string) (authToken string, res *result.Result, err error) {
	return doPost("POST", requestUrl, token, param)
}

// Delete
// 描述：封装https的delete方法
// 创建者：张平
// 创建时间：2022/5/26 17:14
// 参数：
//	 [in]
//		requestUrl：请求的url
//		token：请求携带的token信息
//		param：请求参数
//   [out]
//		authToken：认证token
//		body：请求返回的数据
//		err：请求过程中出现的错误
//   [in/out]
//		无
func Delete(requestUrl, token, param string) (authToken string, res *result.Result, err error) {
	return doPost("DELETE", requestUrl, token, param)
}

// Put
// 描述：封装https的put方法
// 创建者：张平
// 创建时间：2022/7/22 9:51
// 参数：
//	 [in]
//		requestUrl：请求的url
//		token：请求携带的token信息
//		param：请求参数
//   [out]
//		authToken：认证token
//		body：请求返回的数据
//		err：请求过程中出现的错误
//   [in/out]
//		无
func Put(requestUrl, token, param string) (authToken string, res *result.Result, err error) {
	return doPost("PUT", requestUrl, token, param)
}
