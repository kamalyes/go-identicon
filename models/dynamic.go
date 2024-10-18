/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-10-18 11:55:06
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-10-18 12:54:53
 * @FilePath: \go-identicon\models\dynamic.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package models

import (
	"github.com/kamalyes/go-identicon/enums"
	"github.com/kamalyes/go-identicon/types"
)

// 定义一个结构体来存储每个部件的类型和路径
type Widget struct {
	WidgetType types.WidgetType
	Shape      string
}

// 定义一个辅助函数来生成 SVG 路径
func generatePaths(widgets []Widget, prefix string) types.WidgetData {
	widgetData := make(types.WidgetData)

	for _, widget := range widgets {
		if widgetData[widget.WidgetType] == nil {
			widgetData[widget.WidgetType] = make(map[string]types.LoadSVGFunc)
		}
		widgetData[widget.WidgetType][widget.Shape] = func(path string) func() string {
			return func() string { return path }
		}(prefix + string(widget.WidgetType) + "/" + widget.Shape + ".svg")
	}

	return widgetData
}

// 定义一个包含所有部件的列表
var widgets = []Widget{
	{enums.FaceWidget, string(types.FaceShape(enums.BaseFace))},
	{enums.EarWidget, string(types.EarShape(enums.AttachedEar))},
	{enums.EarWidget, string(types.EarShape(enums.DetachedEar))},
	{enums.EyesWidget, string(types.EyesShape(enums.EllipseEyes))},
	{enums.EyesWidget, string(types.EyesShape(enums.EyeshadowEyes))},
	{enums.EyesWidget, string(types.EyesShape(enums.RoundEyes))},
	{enums.EyesWidget, string(types.EyesShape(enums.SmilingEyes))},
	{enums.BeardWidget, string(types.BeardShape(enums.ScruffBeard))},
	{enums.ClothesWidget, string(types.ClothesShape(enums.CollaredClothes))},
	{enums.ClothesWidget, string(types.ClothesShape(enums.CrewClothes))},
	{enums.ClothesWidget, string(types.ClothesShape(enums.OpenClothes))},
	{enums.EarringsWidget, string(types.EarringsShape(enums.HoopEarrings))},
	{enums.EarringsWidget, string(types.EarringsShape(enums.StudEarrings))},
	{enums.EyebrowsWidget, string(types.EyebrowsShape(enums.DownEyebrows))},
	{enums.EyebrowsWidget, string(types.EyebrowsShape(enums.EyelashesDown))},
	{enums.EyebrowsWidget, string(types.EyebrowsShape(enums.EyelashesUp))},
	{enums.EyebrowsWidget, string(types.EyebrowsShape(enums.UpEyebrows))},
	{enums.GlassesWidget, string(types.GlassesShape(enums.RoundGlasses))},
	{enums.GlassesWidget, string(types.GlassesShape(enums.SquareGlasses))},
	{enums.MouthWidget, string(types.MouthShape(enums.FrownMouth))},
	{enums.MouthWidget, string(types.MouthShape(enums.LaughingMouth))},
	{enums.MouthWidget, string(types.MouthShape(enums.NervousMouth))},
	{enums.MouthWidget, string(types.MouthShape(enums.PuckerMouth))},
	{enums.MouthWidget, string(types.MouthShape(enums.SadMouth))},
	{enums.MouthWidget, string(types.MouthShape(enums.SmileMouth))},
	{enums.MouthWidget, string(types.MouthShape(enums.SmirkMouth))},
	{enums.MouthWidget, string(types.MouthShape(enums.SurprisedMouth))},
	{enums.NoseWidget, string(types.NoseShape(enums.CurveNose))},
	{enums.NoseWidget, string(types.NoseShape(enums.PointedNose))},
	{enums.NoseWidget, string(types.NoseShape(enums.RoundNose))},
	{enums.TopsWidget, string(types.TopsShape(enums.BeanieTops))},
	{enums.TopsWidget, string(types.TopsShape(enums.CleanTops))},
	{enums.TopsWidget, string(types.TopsShape(enums.DannyTops))},
	{enums.TopsWidget, string(types.TopsShape(enums.FonzeTops))},
	{enums.TopsWidget, string(types.TopsShape(enums.FunnyTops))},
	{enums.TopsWidget, string(types.TopsShape(enums.PixieTops))},
	{enums.TopsWidget, string(types.TopsShape(enums.PunkTops))},
	{enums.TopsWidget, string(types.TopsShape(enums.TurbanTops))},
	{enums.TopsWidget, string(types.TopsShape(enums.WaveTops))},
}

// 生成 WidgetData
var WidgetData = generatePaths(widgets, "../resources/widgets/")

// 生成 PreviewData
var PreviewData = generatePaths(widgets, "../resources/preview/")
