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
	Shape      types.ShapeStringer
}

// 定义一个辅助函数来生成 SVG 路径
func generatePaths(widgets []Widget, prefix string) types.WidgetData {
	widgetData := make(types.WidgetData)

	for _, widget := range widgets {
		if widgetData[widget.WidgetType] == nil {
			widgetData[widget.WidgetType] = make(map[string]types.LoadSVGFunc)
		}
		widgetData[widget.WidgetType][widget.Shape.String()] = func(path string) func() string {
			return func() string { return path }
		}(prefix + string(widget.WidgetType) + "/" + widget.Shape.String() + ".svg")
	}

	return widgetData
}

// 定义一个包含所有部件的列表
var widgets = []Widget{
	{enums.FaceWidget, enums.BaseFace},
	{enums.EarWidget, enums.AttachedEar},
	{enums.EarWidget, enums.DetachedEar},
	{enums.EyesWidget, enums.EllipseEyes},
	{enums.EyesWidget, enums.EyeshadowEyes},
	{enums.EyesWidget, enums.RoundEyes},
	{enums.EyesWidget, enums.SmilingEyes},
	{enums.BeardWidget, enums.ScruffBeard},
	{enums.ClothesWidget, enums.CollaredClothes},
	{enums.ClothesWidget, enums.CrewClothes},
	{enums.ClothesWidget, enums.OpenClothes},
	{enums.EarringsWidget, enums.HoopEarrings},
	{enums.EarringsWidget, enums.StudEarrings},
	{enums.EyebrowsWidget, enums.DownEyebrows},
	{enums.EyebrowsWidget, enums.EyelashesDown},
	{enums.EyebrowsWidget, enums.EyelashesUp},
	{enums.EyebrowsWidget, enums.UpEyebrows},
	{enums.GlassesWidget, enums.RoundGlasses},
	{enums.GlassesWidget, enums.SquareGlasses},
	{enums.MouthWidget, enums.FrownMouth},
	{enums.MouthWidget, enums.LaughingMouth},
	{enums.MouthWidget, enums.NervousMouth},
	{enums.MouthWidget, enums.PuckerMouth},
	{enums.MouthWidget, enums.SadMouth},
	{enums.MouthWidget, enums.SmileMouth},
	{enums.MouthWidget, enums.SmirkMouth},
	{enums.MouthWidget, enums.SurprisedMouth},
	{enums.NoseWidget, enums.CurveNose},
	{enums.NoseWidget, enums.PointedNose},
	{enums.NoseWidget, enums.RoundNose},
	{enums.TopsWidget, enums.BeanieTops},
	{enums.TopsWidget, enums.CleanTops},
	{enums.TopsWidget, enums.DannyTops},
	{enums.TopsWidget, enums.FonzeTops},
	{enums.TopsWidget, enums.FunnyTops},
	{enums.TopsWidget, enums.PixieTops},
	{enums.TopsWidget, enums.PunkTops},
	{enums.TopsWidget, enums.TurbanTops},
	{enums.TopsWidget, enums.WaveTops},
}

// 生成 WidgetData
var WidgetData = generatePaths(widgets, "../resources/widgets/")

// 生成 PreviewData
var PreviewData = generatePaths(widgets, "../resources/preview/")
