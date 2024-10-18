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
	Gender       Gender
	WrapperShape WrapperShape
	Background   Background
	Widgets      Widgets
}
