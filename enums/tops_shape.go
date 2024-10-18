/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-10-18 10:35:56
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-10-18 11:05:55
 * @FilePath: \go-identicon\enums\tops_shape.go
 * @Description: 定义不同上衣样式的常量
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package enums

import (
	"github.com/kamalyes/go-identicon/types"
)

// 定义上衣样式的常量
const (
	FonzeTops  types.TopsShape = "fonze"  // 冷酷风格
	FunnyTops  types.TopsShape = "funny"  // 有趣风格
	CleanTops  types.TopsShape = "clean"  // 干净风格
	PunkTops   types.TopsShape = "punk"   // 朋克风格
	DannyTops  types.TopsShape = "danny"  // 丹尼风格
	WaveTops   types.TopsShape = "wave"   // 波浪风格
	TurbanTops types.TopsShape = "turban" // 头巾风格
	PixieTops  types.TopsShape = "pixie"  // 小精灵风格
	BeanieTops types.TopsShape = "beanie" // 毛线帽风格
)

// TopsShapes 包含所有上衣样式的数组
var TopsShapes = []types.TopsShape{
	FonzeTops,
	FunnyTops,
	CleanTops,
	PunkTops,
	DannyTops,
	WaveTops,
	TurbanTops,
	PixieTops,
	BeanieTops,
}
