/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-10-18 11:11:19
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-10-18 15:59:47
 * @FilePath: \go-identicon\types\color.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package types

// BackgroundColor 定义背景颜色类型
type BackgroundColor string

// BorderColor 定义边框颜色类型
type BorderColor string

// CommonColors 定义常用颜色类型
type CommonColors string

// SkinColors 定义肤色类型
type SkinColors string

// Color 接口定义了所有颜色类型应实现的方法
type ColorStringer interface {
	String() string
}

// String 方法返回 BackgroundColor 的字符串表示
func (b BackgroundColor) String() string {
	return string(b)
}

// String 方法返回 BorderColor 的字符串表示
func (b BorderColor) String() string {
	return string(b)
}

// String 方法返回 CommonColors 的字符串表示
func (c CommonColors) String() string {
	return string(c)
}

// String 方法返回 SkinColors 的字符串表示
func (s SkinColors) String() string {
	return string(s)
}

// ConvertColorsToStrings 将实现 ColorStringer 接口的切片转换为 []string
func ConvertColorsToStrings(colors []ColorStringer) []string {
	stringColors := make([]string, len(colors))
	for i, color := range colors {
		stringColors[i] = color.String() // 调用 String 方法获取字符串表示
	}
	return stringColors
}

// ConvertToColorStringerSlice 类型转换函数
func ConvertToColorStringerSlice(colors interface{}) []ColorStringer {
	switch v := colors.(type) {
	case []SkinColors:
		result := make([]ColorStringer, len(v))
		for i := range v {
			result[i] = v[i]
		}
		return result
	case []CommonColors:
		result := make([]ColorStringer, len(v))
		for i := range v {
			result[i] = v[i]
		}
		return result
	case []BackgroundColor:
		result := make([]ColorStringer, len(v))
		for i := range v {
			result[i] = v[i]
		}
		return result
	case []BorderColor:
		result := make([]ColorStringer, len(v))
		for i := range v {
			result[i] = v[i]
		}
		return result
	default:
		return nil
	}
}
