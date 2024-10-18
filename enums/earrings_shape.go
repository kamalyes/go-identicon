/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-10-18 10:35:56
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-10-18 13:18:26
 * @FilePath: \go-identicon\enums\earrings_shape.go
 * @Description: 定义不同耳环的常量
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package enums

import (
	"github.com/kamalyes/go-identicon/types"
)

const (
	HoopEarrings types.EarringsShape = "hoop" // 圈形耳环
	StudEarrings types.EarringsShape = "stud" // 钉子耳环
	NoneEarrings types.EarringsShape = "none" // 无耳环
)

// EarringsShapes 包含所有耳环的数组
var EarringsShapes = []types.EarringsShape{
	HoopEarrings,
	StudEarrings,
	NoneEarrings,
}
