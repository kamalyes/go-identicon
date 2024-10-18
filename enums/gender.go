/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-10-18 10:35:56
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-10-18 17:39:54
 * @FilePath: \go-identicon\enums\gender.go
 * @Description: 定义不同性别的常量
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package enums

import (
	"github.com/kamalyes/go-identicon/types"
)

// 定义不同性别的常量
const (
	MaleGender   types.Gender = "male"   // 男性
	FemaleGender types.Gender = "female" // 女性
	NotSetGender types.Gender = "notSet" // 未设定
)

// Genders 包含所有性别的数组
var Genders = []types.Gender{
	MaleGender,
	FemaleGender,
}
