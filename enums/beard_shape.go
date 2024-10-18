/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-10-18 10:35:56
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-10-18 10:56:19
 * @FilePath: \go-identicon\enums\beard_shape.go
 * @Description: 定义不同胡子形状的常量
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package enums

import (
	"github.com/kamalyes/go-identicon/types"
)

const (
	ScruffBeard types.BeardShape = "scruff" // 凌乱胡子
	NoneBeard   types.BeardShape = "none"   // 无胡子
)

// BeardShapes 包含所有胡子形状的数组
var BeardShapes = []types.BeardShape{
	ScruffBeard,
	NoneBeard,
}
