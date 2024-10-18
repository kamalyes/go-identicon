/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-10-18 11:58:55
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-10-18 16:31:05
 * @FilePath: \go-identicon\types\avatar.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package types

// AvatarWidgets 结构体包含多个小部件，用于构建头像
type AvatarWidgets struct {
	Face     Widget // 脸部小部件
	Tops     Widget // 上衣小部件
	Ear      Widget // 耳朵小部件
	Earrings Widget // 耳环小部件
	Eyebrows Widget // 眉毛小部件
	Glasses  Widget // 眼镜小部件
	Eyes     Widget // 眼睛小部件
	Nose     Widget // 鼻子小部件
	Mouth    Widget // 嘴巴小部件
	Beard    Widget // 胡子小部件
	Clothes  Widget // 衣服小部件
}

// Background 结构体表示头像的背景设置
type Background struct {
	BackgroundColor BackgroundColor // 背景颜色
	BorderColor     BorderColor     // 边框颜色
}

// AvatarSettings 结构体包含头像的所有设置选项
type AvatarSettings struct {
	Genders          []Gender          // 性别选项数组，支持多种性别
	WrapperShapes    []WrapperShape    // 可选的包装形状数组
	FaceShapes       []FaceShape       // 可选的脸部形状数组
	TopsShapes       []TopsShape       // 可选的上衣形状数组
	EarShapes        []EarShape        // 可选的耳朵形状数组
	EarringsShapes   []EarringsShape   // 可选的耳环形状数组
	EyebrowsShapes   []EyebrowsShape   // 可选的眉毛形状数组
	EyesShapes       []EyesShape       // 可选的眼睛形状数组
	NoseShapes       []NoseShape       // 可选的鼻子形状数组
	MouthShapes      []MouthShape      // 可选的嘴巴形状数组
	BeardShapes      []BeardShape      // 可选的胡子形状数组
	GlassesShapes    []GlassesShape    // 可选的眼镜形状数组
	ClothesShapes    []ClothesShape    // 可选的衣服形状数组
	CommonColors     []CommonColors    // 常用颜色数组
	SkinColors       []SkinColors      // 皮肤颜色数组
	BackgroundColors []BackgroundColor // 背景颜色数组
	BorderColors     []BorderColor     // 边框颜色数组
}

// AvatarOption 定义了头像选项的结构体
type AvatarOption struct {
	Gender       Gender
	WrapperShape WrapperShape
	Background   Background
	Widgets      struct {
		Beard struct {
			Shape     BeardShape
			FillColor CommonColors
		}
		Face struct {
			Shape     FaceShape
			FillColor CommonColors
		}
		Tops struct {
			Shape     TopsShape
			FillColor CommonColors
		}
		Ear struct {
			Shape EarShape
		}
		Earrings struct {
			Shape EarringsShape
		}
		Eyebrows struct {
			Shape EyebrowsShape
		}
		Eyes struct {
			Shape EyesShape
		}
		Nose struct {
			Shape NoseShape
		}
		Glasses struct {
			Shape GlassesShape
		}
		Mouth struct {
			Shape MouthShape
		}
		Clothes struct {
			Shape     ClothesShape
			FillColor CommonColors
		}
	}
}
