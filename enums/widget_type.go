/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-10-18 10:35:56
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-10-18 11:26:01
 * @FilePath: \go-identicon\enums\widget_type.go
 * @Description: 定义与小部件类型相关的常量
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package enums

import (
	"github.com/kamalyes/go-identicon/types"
)

const (
	FaceWidget     types.WidgetType = "face"     // 脸
	TopsWidget     types.WidgetType = "tops"     // 上衣
	EarWidget      types.WidgetType = "ear"      // 耳朵
	EarringsWidget types.WidgetType = "earrings" // 耳环
	EyebrowsWidget types.WidgetType = "eyebrows" // 眉毛
	EyesWidget     types.WidgetType = "eyes"     // 眼睛
	NoseWidget     types.WidgetType = "nose"     // 鼻子
	GlassesWidget  types.WidgetType = "glasses"  // 眼镜
	MouthWidget    types.WidgetType = "mouth"    // 嘴巴
	BeardWidget    types.WidgetType = "beard"    // 胡子
	ClothesWidget  types.WidgetType = "clothes"  // 衣服
)

// Widgets 包含所有小部件类型的数组
var Widgets = []types.WidgetType{
	FaceWidget,
	TopsWidget,
	EarWidget,
	EarringsWidget,
	EyebrowsWidget,
	EyesWidget,
	NoseWidget,
	GlassesWidget,
	MouthWidget,
	BeardWidget,
	ClothesWidget,
}
