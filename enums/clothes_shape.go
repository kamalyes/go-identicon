/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-10-18 10:35:56
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-10-18 11:08:25
 * @FilePath: \go-identicon\enums\clothes_shape.go
 * @Description: 定义不同衣服领子的常量
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package enums

import (
	"github.com/kamalyes/go-identicon/types"
)

const (
	CrewClothes     types.ClothesShape = "crew"     // 圆领
	CollaredClothes types.ClothesShape = "collared" // 有领
	OpenClothes     types.ClothesShape = "open"     // 开领
)

// ClothesShapes 包含所有衣服领子的数组
var ClothesShapes = []types.ClothesShape{
	CrewClothes,
	CollaredClothes,
	OpenClothes,
}
