/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-10-18 10:35:56
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-10-18 13:45:25
 * @FilePath: \go-identicon\enums\wrapper_shape.go
 * @Description: 定义与包装形状相关的常量
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package enums

import (
	"github.com/kamalyes/go-identicon/types"
)

const (
	CircleWrapper   types.WrapperShape = "circle"   // 圆形
	SquareWrapper   types.WrapperShape = "square"   // 正方形
	SquircleWrapper types.WrapperShape = "squircle" // 圆角方形
)

// WrapperShapes 包含所有包装形状的数组
var WrapperShapes = []types.WrapperShape{
	CircleWrapper,
	SquareWrapper,
	SquircleWrapper,
}
