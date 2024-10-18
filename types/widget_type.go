/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-10-18 11:11:19
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-10-19 13:01:56
 * @FilePath: \go-identicon\types\widget_type.go
 * @Description: 定义形状类型
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package types

import "fmt"

// WidgetStringer 接口定义了 String 方法
type WidgetStringer interface {
	String() string
}

// Widget 结构体表示一个小部件
type Widget struct {
	Shape       interface{}   // 小部件的形状，可以是任意类型或 None
	ZIndex      int           // Z轴索引，使用指针以支持可选字段
	FillColor   ColorStringer // 填充颜色，使用指针以支持可选字段
	StrokeColor string        // 描边颜色，使用指针以支持可选字段
}

// Widgets 定义组件结构
type Widgets struct {
	Face     Widget
	Tops     Widget
	Ear      Widget
	Earrings Widget
	Eyebrows Widget
	Eyes     Widget
	Nose     Widget
	Glasses  Widget
	Mouth    Widget
	Beard    Widget
	Clothes  Widget
}

// WidgetBase 结构体表示小部件的基础部分
type WidgetBase struct {
	FillColor CommonColors
}

// BeardWidget 结构体表示胡子组件
type BeardWidget struct {
	WidgetBase
	Shape BeardShape
}

// FaceWidget 结构体表示脸部组件
type FaceWidget struct {
	WidgetBase
	Shape FaceShape
}

// TopsWidget 结构体表示上衣组件
type TopsWidget struct {
	WidgetBase
	Shape TopsShape
}

// EarWidget 结构体表示耳朵组件
type EarWidget struct {
	WidgetBase
	Shape EarShape
}

// EarringsWidget 结构体表示耳环组件
type EarringsWidget struct {
	WidgetBase
	Shape EarringsShape
}

// EyebrowsWidget 结构体表示眉毛组件
type EyebrowsWidget struct {
	WidgetBase
	Shape EyebrowsShape
}

// EyesWidget 结构体表示眼睛组件
type EyesWidget struct {
	WidgetBase
	Shape EyesShape
}

// NoseWidget 结构体表示鼻子组件
type NoseWidget struct {
	WidgetBase
	Shape NoseShape
}

// GlassesWidget 结构体表示眼镜组件
type GlassesWidget struct {
	WidgetBase
	Shape GlassesShape
}

// MouthWidget 结构体表示嘴巴组件
type MouthWidget struct {
	WidgetBase
	Shape MouthShape
}

// ClothesWidget 结构体表示衣服组件
type ClothesWidget struct {
	WidgetBase
	Shape ClothesShape
}

// 实现 BeardWidget 的 String 方法
func (b BeardWidget) String() string {
	return fmt.Sprintf("Beard: Shape=%s, FillColor=%s", b.Shape, b.FillColor)
}

// 实现 FaceWidget 的 String 方法
func (f FaceWidget) String() string {
	return fmt.Sprintf("Face: Shape=%s, FillColor=%s", f.Shape, f.FillColor)
}

// 实现 TopsWidget 的 String 方法
func (t TopsWidget) String() string {
	return fmt.Sprintf("Tops: Shape=%s, FillColor=%s", t.Shape, t.FillColor)
}

// 实现 EarWidget 的 String 方法
func (e EarWidget) String() string {
	return fmt.Sprintf("Ear: Shape=%s", e.Shape)
}

// 实现 EarringsWidget 的 String 方法
func (e EarringsWidget) String() string {
	return fmt.Sprintf("Earrings: Shape=%s", e.Shape)
}

// 实现 EyebrowsWidget 的 String 方法
func (e EyebrowsWidget) String() string {
	return fmt.Sprintf("Eyebrows: Shape=%s", e.Shape)
}

// 实现 EyesWidget 的 String 方法
func (e EyesWidget) String() string {
	return fmt.Sprintf("Eyes: Shape=%s", e.Shape)
}

// 实现 NoseWidget 的 String 方法
func (n NoseWidget) String() string {
	return fmt.Sprintf("Nose: Shape=%s", n.Shape)
}

// 实现 GlassesWidget 的 String 方法
func (g GlassesWidget) String() string {
	return fmt.Sprintf("Glasses: Shape=%s", g.Shape)
}

// 实现 MouthWidget 的 String 方法
func (m MouthWidget) String() string {
	return fmt.Sprintf("Mouth: Shape=%s", m.Shape)
}

// 实现 ClothesWidget 的 String 方法
func (c ClothesWidget) String() string {
	return fmt.Sprintf("Clothes: Shape=%s, FillColor=%s", c.Shape, c.FillColor)
}
