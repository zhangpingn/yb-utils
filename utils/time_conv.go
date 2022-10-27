// Copyright (c) 2023 YaoBase
// 描述: 封装日期格式转化相关方法
// 创建者: 张平
// 创建时间: 2022/5/26 17:13

package utils

import (
	"time"
)

// TimeFromInt64
// 描述：时间戳转换为时间对象实例
// 创建者：张平
// 创建时间：2022/6/13 17:14
// 参数：
//	 [in]
//		n：时间戳
//   [out]
//		时间对象实例
//   [in/out]
//		无
func TimeFromInt64(n int64) time.Time {
	return time.Unix(n, 0)
}

// TimeFromMill
// 描述：毫秒时间戳转换为时间对象实例
// 创建者：张平
// 创建时间：2022/6/13 17:14
// 参数：
//	 [in]
//		n：时间戳
//   [out]
//		时间对象实例
//   [in/out]
//		无
func TimeFromMill(n int64) time.Time {
	return time.UnixMilli(n)
}

// TimeToString
// 描述：时间对象实例格式化成字符串
// 创建者：张平
// 创建时间：2022/6/13 17:14
// 参数：
//	 [in]
//		t：时间对象实例
//   [out]
//		格式化后的字符串
//   [in/out]
//		无
func TimeToString(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// TimeFromString
// 描述：字符串转换为时间对象实例
// 创建者：张平
// 创建时间：2022/6/13 17:14
// 参数：
//	 [in]
//		s：待转化的字符串
//   [out]
//		结果1：时间对象实例
//		结果2：转化过程中出现的错误
//   [in/out]
//		无
func TimeFromString(s string) (time.Time, error) {
	return time.ParseInLocation("2006-01-02 15:04:05", s, time.Local)
}
