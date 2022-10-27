// Copyright (c) 2023 YaoBase
// 描述: 定义前端页面分页
// 创建者: 张平
// 创建时间: 2022/6/13 17:12

package result

import (
	"YBFacadeService/common"
)

// PageData 分页模型，用于进行传递分页功能的参数
type PageData struct {
	PageNum  int64 `form:"page_num" json:"page_num"`   // 当前页数
	PageSize int64 `form:"page_size" json:"page_size"` // 每页多少条数据
}

// GetPage
// 描述：分页操作，进行分页变量值的设置
// 创建者：张平
// 创建时间：2022/6/13 17:13
// 参数：
//	 [in]
//		flag：分页类型
//		count：当前处理数据总量
//   [out]
//		无
//   [in/out]
//		p：分页对象实例
func (p *PageData) GetPage(flag int, count ...int64) {
	if p.PageNum <= 0 {
		p.PageNum = 1
	}
	if p.PageSize <= 0 {
		p.PageSize = 10 // 默认每页显示10条数据
	}
	start := p.PageNum
	switch flag {
	case common.REDIS:
		p.PageNum = (p.PageNum - 1) * p.PageSize
		// redis的lrange是包含尾端
		if p.PageSize == 1 {
			p.PageSize = (start-1)*p.PageSize + p.PageSize
		} else {
			p.PageSize = (start-1)*p.PageSize + p.PageSize - 1
		}
	case common.MYSQL:
		if len(count) > 0 {
			// 当总条数小于当前分页的数量时，应该返回所有数据，及limit 0,总条数
			if p.PageSize >= count[0] {
				p.PageNum = 0
				p.PageSize = count[0]
			} else {
				s := count[0] / p.PageSize
				y := count[0] % p.PageSize

				// 总条数除以每页的数量求余大于0，故其总页数应加一
				if y > 0 {
					s++
				}
				// 当前页数若大于总页数-1，则标书其带的当前页数有误，故在此给其最后一组，及总页数
				if p.PageNum > s-1 {
					p.PageNum = s
				}
			}
			if p.PageNum > 0 {
				p.PageNum = (p.PageNum - 1) * p.PageSize
			}
		}
	}
}
