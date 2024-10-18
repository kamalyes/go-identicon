/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-10-18 11:11:19
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-10-18 19:26:24
 * @FilePath: \go-identicon\types\shape.go
 * @Description: 定义形状类型
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package types

// BeardShape 定义了胡子形状的常量
type BeardShape string

// ClothesShape 定义了衣服形状的常量
type ClothesShape string

// EarringsShape 定义了耳环形状的常量
type EarringsShape string

// EarShape 定义了耳朵形状的常量
type EarShape string

// EyebrowsShape 定义了眉毛形状的常量
type EyebrowsShape string

// EyesShape 定义了眼睛形状的常量
type EyesShape string

// FaceShape 定义了脸部形状的常量
type FaceShape string

// GlassesShape 定义了眼镜形状的常量
type GlassesShape string

// MouthShape 定义了嘴巴形状的常量
type MouthShape string

// NoseShape 定义了鼻子形状的常量
type NoseShape string

// TopsShape 定义了上衣形状的常量
type TopsShape string

// WrapperShape 定义了包装形状的常量
type WrapperShape string

// WidgetShape 定义了小部件形状的类型，可以是任意形状
type WidgetShape interface{}

// ShapeStringer 接口定义了所有形状类型应实现的方法
type ShapeStringer interface {
	String() string
}

// String 方法返回 Gender 的字符串表示
func (g Gender) String() string {
	return string(g)
}

// String 方法返回 WidgetType 的字符串表示
func (w WidgetType) String() string {
	return string(w)
}

// String 方法返回 ActionType 的字符串表示
func (a ActionType) String() string {
	return string(a)
}

// String 方法返回 BeardShape 的字符串表示
func (b BeardShape) String() string {
	return string(b)
}

// String 方法返回 ClothesShape 的字符串表示
func (c ClothesShape) String() string {
	return string(c)
}

// String 方法返回 EarringsShape 的字符串表示
func (e EarringsShape) String() string {
	return string(e)
}

// String 方法返回 EarShape 的字符串表示
func (e EarShape) String() string {
	return string(e)
}

// String 方法返回 EyebrowsShape 的字符串表示
func (e EyebrowsShape) String() string {
	return string(e)
}

// String 方法返回 EyesShape 的字符串表示
func (e EyesShape) String() string {
	return string(e)
}

// String 方法返回 FaceShape 的字符串表示
func (f FaceShape) String() string {
	return string(f)
}

// String 方法返回 GlassesShape 的字符串表示
func (g GlassesShape) String() string {
	return string(g)
}

// String 方法返回 MouthShape 的字符串表示
func (m MouthShape) String() string {
	return string(m)
}

// String 方法返回 NoseShape 的字符串表示
func (n NoseShape) String() string {
	return string(n)
}

// String 方法返回 TopsShape 的字符串表示
func (t TopsShape) String() string {
	return string(t)
}

// String 方法返回 WrapperShape 的字符串表示
func (w WrapperShape) String() string {
	return string(w)
}
