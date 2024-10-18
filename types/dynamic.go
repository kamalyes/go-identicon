/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-10-18 11:55:06
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-10-18 13:16:50
 * @FilePath: \go-identicon\models\avatar.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package types

// 定义一个类型用于加载 SVG 图像
type LoadSVGFunc func() string

// 定义一个结构体来存储每个部件的类型和路径
type WidgetShapeMap struct {
	WidgetType WidgetType
	Shape      string
}

// 定义一个数据结构来存储 widgetData 和 previewData
type WidgetData map[WidgetType]map[string]LoadSVGFunc
