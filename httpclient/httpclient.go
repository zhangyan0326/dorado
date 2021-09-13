package httpclient

import (
	"bytes"
	"dorado/bizerror"
	"dorado/zlog"
	"github.com/kataras/iris/v12"
	"io/ioutil"
	"net/http"
	"time"
)

func Get(ctx *iris.Context, url string) {

}

func GetWithParams(ctx *iris.Context, url string, params map[string]interface{}) {

}

/**
Post
*/
func Post(authorization string, url string, method string, jsonStr string) (string, string, *bizerror.BizError) {
	zlog.Info("【Post】请求进入网络请求 Start...", jsonStr)

	req, errRequest := http.NewRequest(method, url, bytes.NewBuffer([]byte(jsonStr)))
	if errRequest != nil {
		zlog.Error("【Post】组装请求失败")
		return "", "", &bizerror.HttpError
	}
	req.Header.Add("Content-type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", authorization)

	defer req.Body.Close()
	client := &http.Client{Timeout: 10 * time.Second}
	resp, errClient := client.Do(req)
	zlog.Info("【Post】请求客户端是:", req)
	if errClient != nil {
		zlog.Error("【Post】请求http失败:", errClient)
		return "", "", &bizerror.HttpError
	}
	result, errReade := ioutil.ReadAll(resp.Body)
	status := resp.Status
	if errReade != nil {
		zlog.Error("【Post】请求失败:", errClient)
		return "", "", &bizerror.HttpError
	}
	zlog.Info("【Post】请求进入网络请求 End...", string(result))
	return string(result), status, nil
}

func PostLeo(authorization string, url string, method string, jsonStr string) (string, string, *bizerror.BizError) {
	zlog.Info("【Post】请求进入网络请求 Start...", jsonStr)
	req, errRequest := http.NewRequest(method, url, bytes.NewBuffer([]byte(jsonStr)))
	if errRequest != nil {
		zlog.Error("【Post】组装请求失败")
		return "", "", &bizerror.HttpError
	}
	req.Header.Add("Content-type", "application/json")
	req.Header.Add("Authorization", authorization)
	defer req.Body.Close()
	client := &http.Client{Timeout: 10 * time.Second}
	resp, errClient := client.Do(req)
	zlog.Info("【Post】请求客户端是:", req)
	if errClient != nil {
		zlog.Error("【Post】请求http失败:", errClient)
		return "", "", &bizerror.HttpError
	}
	result, errReade := ioutil.ReadAll(resp.Body)
	status := resp.Status
	if errReade != nil {
		zlog.Error("【Post】请求失败:", errClient)
		return "", "", &bizerror.HttpError
	}
	zlog.Info("【Post】请求进入网络请求 End...", string(result))
	return string(result), status, nil
}
