/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-10-18 10:35:56
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-10-18 10:15:51
 * @FilePath: \go-identicon\enums\action_type.go
 * @Description: 定义不同嘴型的常量
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package enums

import (
	"github.com/kamalyes/go-identicon/types"
)

// 定义不同嘴型的常量
const (
	FrownMouth     types.MouthShape = "frown"     // 皱眉
	LaughingMouth  types.MouthShape = "laughing"  // 笑
	NervousMouth   types.MouthShape = "nervous"   // 紧张
	PuckerMouth    types.MouthShape = "pucker"    // 嘴唇嘟起
	SadMouth       types.MouthShape = "sad"       // 悲伤
	SmileMouth     types.MouthShape = "smile"     // 微笑
	SmirkMouth     types.MouthShape = "smirk"     // 笑而不语
	SurprisedMouth types.MouthShape = "surprised" // 惊讶
)

// MouthShapes 包含所有嘴型的数组
var MouthShapes = []types.MouthShape{
	FrownMouth,
	LaughingMouth,
	NervousMouth,
	PuckerMouth,
	SadMouth,
	SmileMouth,
	SmirkMouth,
	SurprisedMouth,
}
