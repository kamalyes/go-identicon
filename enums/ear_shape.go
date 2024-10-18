/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-10-18 10:35:56
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-10-18 10:50:16
 * @FilePath: \go-identicon\enums\ear_shape.go
 * @Description: 定义不同耳朵的常量
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package enums

import (
	"github.com/kamalyes/go-identicon/types"
)

const (
	AttachedEar types.EarShape = "attached" // 附耳
	DetachedEar types.EarShape = "detached" // 自由耳
)

// EarShapes 包含所有耳朵的数组
var EarShapes = []types.EarShape{
	AttachedEar,
	DetachedEar,
}
