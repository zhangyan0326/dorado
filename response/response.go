package response

import (
	"dorado/bizerror"
	"encoding/json"
	"github.com/kataras/iris/v12"
	log "github.com/sirupsen/logrus"
	"time"
)

type CommonResponse struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Error(ctx iris.Context, bizError *bizerror.BizError) {
	responseCommon := CommonResponse{
		Code: bizError.Code,
		Msg:  bizError.Message,
		Data: nil,
	}
	ctx.JSON(responseCommon)
	setMonitorValue(ctx, responseCommon)
	log.Info("【CommonResponseEnd】返回客户端的响应 :", responseCommon.Code, responseCommon.Msg)
	return
}

func Success(ctx iris.Context, data interface{}) {
	responseCommon := CommonResponse{
		Code: "0000",
		Msg:  "success",
		Data: data,
	}
	ctx.JSON(responseCommon)
	setMonitorValue(ctx, responseCommon)
	log.Info("【CommonResponseEnd】返回客户端的响应 :", responseCommon.Code, responseCommon.Msg)
	return
}

func CommonResponseEnd(ctx iris.Context, response *CommonResponse, bizError *bizerror.BizError) {
	responseCommon := CommonResponse{
		Code: bizError.Code,
		Msg:  bizError.Message,
		Data: nil,
	}
	ctx.JSON(responseCommon)
	setMonitorValue(ctx, responseCommon)
	log.Info("【CommonResponseEnd】返回客户端的响应 :", responseCommon.Code, responseCommon.Msg)
	return
}

// 设置监控信息
func setMonitorValue(ctx iris.Context, responseCommon CommonResponse) {
	responseInfo, _ := json.Marshal(responseCommon)
	ctx.Values().Set("response", responseInfo)
	ctx.Values().Set("code", responseCommon.Code)
	ctx.Values().Set("response_time", time.Now().Format("2006-01-02 15:04:05"))
}

func Ok(ctx iris.Context) {
	responseCommon := CommonResponse{
		Data: "成功",
	}
	ctx.JSON(responseCommon)
	ctx.Next()
}
