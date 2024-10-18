/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-10-18 11:57:29
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-10-18 11:58:15
 * @FilePath: \go-identicon\models\special_avatars.go
 * @Description:
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package models

import (
	"github.com/kamalyes/go-identicon/enums"
	"github.com/kamalyes/go-identicon/types"
)

// SHAPE_STYLE_SET 定义了形状样式设置
var SHAPE_STYLE_SET = map[types.WrapperShape]struct {
	BorderRadius string
}{
	types.WrapperShape(enums.CircleWrapper):   {BorderRadius: "50%"},
	types.WrapperShape(enums.SquareWrapper):   {BorderRadius: "0"},
	types.WrapperShape(enums.SquircleWrapper): {BorderRadius: "25px"},
}
