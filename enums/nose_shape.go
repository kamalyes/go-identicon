/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-10-18 10:35:56
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-10-18 10:51:05
 * @FilePath: \go-identicon\enums\nose_shape.go
 * @Description: 定义不同鼻型的常量
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package enums

import (
	"github.com/kamalyes/go-identicon/types"
)

// 定义不同鼻型的常量
const (
	CurveNose   types.NoseShape = "curve"   // 曲线鼻
	RoundNose   types.NoseShape = "round"   // 圆鼻
	PointedNose types.NoseShape = "pointed" // 尖鼻
)

// NoseShapes 包含所有鼻型的数组
var NoseShapes = []types.NoseShape{
	CurveNose,
	RoundNose,
	PointedNose,
}
