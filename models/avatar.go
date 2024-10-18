/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-10-18 11:55:06
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-10-18 17:18:13
 * @FilePath: \go-identicon\models\avatar.go
 * @Description: 定义与小部件和头像相关的结构体
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package models

import (
	"github.com/kamalyes/go-identicon/constants"
	"github.com/kamalyes/go-identicon/enums"
	"github.com/kamalyes/go-identicon/types"
)

// 定义 AvatarLayer 结构体
type AvatarLayer struct {
	ZIndex int
}

// 定义 AVATAR_LAYER 常量
var AVATAR_LAYER = map[types.WidgetType]AvatarLayer{
	enums.FaceWidget:     {ZIndex: 10},
	enums.EarWidget:      {ZIndex: 102},
	enums.EarringsWidget: {ZIndex: 103},
	enums.EyebrowsWidget: {ZIndex: 70},
	enums.EyesWidget:     {ZIndex: 50},
	enums.NoseWidget:     {ZIndex: 60},
	enums.GlassesWidget:  {ZIndex: 90},
	enums.MouthWidget:    {ZIndex: 100},
	enums.BeardWidget:    {ZIndex: 105},
	enums.TopsWidget:     {ZIndex: 80},
	enums.ClothesWidget:  {ZIndex: 110},
}

// defaultAvatarColors 定义了头像的默认设置
var defaultAvatarColors = types.AvatarSettings{
	CommonColors: []types.CommonColors{
		constants.ColorLightBlue,
		constants.ColorPink,
		constants.ColorYellow,
		constants.ColorLightPurple,
		constants.ColorTurquoise,
		constants.ColorLightPink,
		constants.ColorLightYellow,
		constants.ColorBlue,
		constants.ColorOrange,
		constants.ColorTeal,
		constants.ColorLavender,
		constants.ColorRed,
	},
	SkinColors: []types.SkinColors{
		constants.ColorSkinLight1,
		constants.ColorSkinLight2,
		constants.ColorSkinMedium1,
		constants.ColorSkinMedium2,
		constants.ColorSkinDark,
	},
	BackgroundColors: []types.BackgroundColor{
		constants.ColorLightBlue,
		constants.ColorPink,
		constants.ColorYellow,
		constants.ColorLightPurple,
		constants.ColorGradient1,
		constants.ColorGradient2,
		constants.ColorGradient3,
		constants.ColorGradient4,
		constants.ColorGradient5,
		constants.ColorTransparent,
	},
}

// AvatarSettings 包含所有的常量设置
var AvatarSettings = types.AvatarSettings{
	Genders:          enums.Genders,
	WrapperShapes:    enums.WrapperShapes,
	FaceShapes:       []types.FaceShape{enums.BaseFace},
	TopsShapes:       enums.TopsShapes,
	EarShapes:        enums.EarShapes,
	EarringsShapes:   enums.EarringsShapes,
	EyebrowsShapes:   enums.EyebrowsShapes,
	EyesShapes:       enums.EyesShapes,
	NoseShapes:       enums.NoseShapes,
	MouthShapes:      enums.MouthShapes,
	BeardShapes:      enums.BeardShapes,
	GlassesShapes:    enums.GlassesShapes,
	ClothesShapes:    enums.ClothesShapes,
	CommonColors:     defaultAvatarColors.CommonColors,
	SkinColors:       defaultAvatarColors.SkinColors,
	BackgroundColors: defaultAvatarColors.BackgroundColors,
	BorderColors:     defaultAvatarColors.BorderColors,
}
