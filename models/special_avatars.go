/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-10-18 11:57:29
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-10-18 13:10:26
 * @FilePath: \go-identicon\models\special_avatars.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package models

import (
	"github.com/kamalyes/go-identicon/constants"
	"github.com/kamalyes/go-identicon/enums"
	"github.com/kamalyes/go-identicon/types"
)

// SPECIAL_AVATARS 定义特殊头像的设置
var SPECIAL_AVATARS = []types.SpecialAvatar{
	{
		WrapperShape: "squircle",
		Background: types.Background{
			BackgroundColor: constants.ColorGradient2,
			BorderColor:     constants.ColorTransparent,
		},
		Widgets: types.Widgets{
			Face:     types.Widget{Shape: types.FaceShape(enums.BaseFace), FillColor: constants.ColorSkinLight2},
			Tops:     types.Widget{Shape: types.TopsShape(enums.PixieTops), FillColor: constants.ColorTurquoise},
			Ear:      types.Widget{Shape: types.EarShape(enums.AttachedEar)},
			Earrings: types.Widget{Shape: types.EarringsShape(enums.StudEarrings)},
			Eyebrows: types.Widget{Shape: types.EyebrowsShape(enums.UpEyebrows)},
			Eyes:     types.Widget{Shape: types.EyesShape(enums.EyeshadowEyes)},
			Nose:     types.Widget{Shape: types.NoseShape(enums.PointedNose)},
			Glasses:  types.Widget{Shape: types.GlassesShape(enums.NoneGlasses)},
			Mouth:    types.Widget{Shape: types.MouthShape(enums.LaughingMouth)},
			Beard:    types.Widget{Shape: types.BeardShape(enums.NoneBeard)},
			Clothes:  types.Widget{Shape: types.ClothesShape(enums.CrewClothes), FillColor: constants.ColorLightPurple},
		},
	},
	// 其他特殊头像...
}
