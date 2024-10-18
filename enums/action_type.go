/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-10-18 10:35:56
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-10-18 11:55:22
 * @FilePath: \go-identicon\enums\action_type.go
 * @Description: 定义不同动作类型的常量
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package enums

import (
	"github.com/kamalyes/go-identicon/types"
)

const (
	Flip types.ActionType = "flip" // 翻转
	Code types.ActionType = "code" // 代码
)

// ActionTypes 包含所有动作类型的数组
var ActionTypes = []types.ActionType{
	Flip,
	Code,
}
