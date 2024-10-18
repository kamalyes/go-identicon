/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-10-18 10:35:56
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-10-18 13:00:15
 * @FilePath: \go-identicon\enums\glasses_shape.go
 * @Description: 定义不同眼镜的常量
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package enums

import (
	"github.com/kamalyes/go-identicon/types"
)

// 定义不同眼镜的常量
const (
	RoundGlasses  types.GlassesShape = "round"  // 圆形眼镜
	SquareGlasses types.GlassesShape = "square" // 方形眼镜
	NoneGlasses   types.GlassesShape = "none"   // 无眼镜
)

// GlassesShapes 包含所有眼镜的数组
var GlassesShapes = []types.GlassesShape{
	RoundGlasses,
	SquareGlasses,
	NoneGlasses,
}
