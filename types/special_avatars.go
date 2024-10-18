/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-10-18 11:58:55
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-10-18 11:59:07
 * @FilePath: \go-identicon\types\avatar.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package types

// SpecialAvatar 定义特殊头像的结构
type SpecialAvatar struct {
	WrapperShape string
	Background   Background
	Widgets      Widgets
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
