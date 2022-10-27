// Copyright (c) 2023 YaoBase
// 描述: 定义与前端进行交互返回的消息结构体
// 创建者: 张平
// 创建时间: 2022/6/13 17:12

package result

// Result 封装返回消息体
type Result struct {
	Code    int         `json:"code"`    //状态码
	Count   int         `json:"count"`   //数据总条数
	Message string      `json:"message"` //消息
	Data    interface{} `json:"data"`    //数据
}

// ReturnNoData
// 描述：返回不带数据的消息结构体对象实例
// 创建者：张平
// 创建时间：2022/6/13 17:13
// 参数：
//	 [in]
//		code：响应码
//		message：响应提示信息
//   [out]
//		响应对象实例
//   [in/out]
//		无
func ReturnNoData(code int, message string) *Result {
	return &Result{
		Code:    code,
		Message: message,
	}
}

// ReturnWithData
// 描述：返回带数据的消息结构体对象实例
// 创建者：张平
// 创建时间：2022/6/13 17:13
// 参数：
//	 [in]
//		code：响应码
//		message：响应提示信息
//		data：响应数据
//   [out]
//		响应对象实例
//   [in/out]
//		无
func ReturnWithData(code, count int, message string, data interface{}) *Result {
	return &Result{
		Code:    code,
		Count:   count,
		Message: message,
		Data:    data,
	}
}
