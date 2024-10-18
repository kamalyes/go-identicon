/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-10-18 10:35:56
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-10-18 10:50:58
 * @FilePath: \go-identicon\enums\eyes_shape.go
 * @Description: 定义不同眼睛的常量
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package enums

import (
	"github.com/kamalyes/go-identicon/types"
)

// 定义不同眼睛的常量
const (
	EllipseEyes   types.EyesShape = "ellipse"   // 椭圆形
	SmilingEyes   types.EyesShape = "smiling"   // 微笑
	EyeshadowEyes types.EyesShape = "eyeshadow" // 眼影
	RoundEyes     types.EyesShape = "round"     // 圆形
)

// EyesShapes 包含所有眼睛形状的数组
var EyesShapes = []types.EyesShape{
	EllipseEyes,
	SmilingEyes,
	EyeshadowEyes,
	RoundEyes,
}
