/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-10-18 10:35:56
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-10-18 11:42:48
 * @FilePath: \go-identicon\enums\eyebrows_shape.go
 * @Description: 定义不同眉毛形状的常量
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package enums

import (
	"github.com/kamalyes/go-identicon/types"
)

const (
	UpEyebrows    types.EyebrowsShape = "up"            // 向上
	DownEyebrows  types.EyebrowsShape = "down"          // 向下
	EyelashesUp   types.EyebrowsShape = "eyelashesup"   // 睫毛上
	EyelashesDown types.EyebrowsShape = "eyelashesdown" // 睫毛下
)

// EyebrowsShapes 包含所有眉毛形状的数组
var EyebrowsShapes = []types.EyebrowsShape{
	UpEyebrows,
	DownEyebrows,
	EyelashesUp,
	EyelashesDown,
}
