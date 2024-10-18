/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2024-10-18 10:35:56
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2024-10-18 11:20:15
 * @FilePath: \go-identicon\enums\widget_shape.go
 * @Description: 定义所有小部件形状的类型
 *
 * Copyright (c) 2024 by kamalyes, All Rights Reserved.
 */
package enums

import (
	"github.com/kamalyes/go-identicon/types"
)

// 定义所有小部件形状的切片
var WidgetShapes = []types.WidgetShape{
	types.FaceShape(BaseFace),           // 基础脸型
	types.TopsShape(FonzeTops),          // Fonze 风格上衣
	types.TopsShape(FunnyTops),          // 有趣风格上衣
	types.TopsShape(CleanTops),          // 干净风格上衣
	types.TopsShape(PunkTops),           // 朋克风格上衣
	types.TopsShape(DannyTops),          // Danny 风格上衣
	types.TopsShape(WaveTops),           // 波浪风格上衣
	types.TopsShape(TurbanTops),         // 头巾风格上衣
	types.TopsShape(PixieTops),          // 小精灵风格上衣
	types.TopsShape(BeanieTops),         // 毛线帽风格上衣
	types.EarShape(AttachedEar),         // 附着耳朵
	types.EarShape(DetachedEar),         // 分离耳朵
	types.EarringsShape(HoopEarrings),   // 圈状耳环
	types.EarringsShape(StudEarrings),   // 钉状耳环
	types.EarringsShape(NoneEarrings),   // 无耳环
	types.EyebrowsShape(UpEyebrows),     // 上翘眉毛
	types.EyebrowsShape(DownEyebrows),   // 下垂眉毛
	types.EyebrowsShape(EyelashesUp),    // 上翘睫毛
	types.EyebrowsShape(EyelashesDown),  // 下垂睫毛
	types.EyesShape(EllipseEyes),        // 椭圆形眼睛
	types.EyesShape(SmilingEyes),        // 微笑眼
	types.EyesShape(EyeshadowEyes),      // 眼影眼睛
	types.EyesShape(RoundEyes),          // 圆形眼睛
	types.NoseShape(CurveNose),          // 曲线鼻子
	types.NoseShape(RoundNose),          // 圆形鼻子
	types.NoseShape(PointedNose),        // 尖鼻子
	types.MouthShape(FrownMouth),        // 皱眉嘴巴
	types.MouthShape(LaughingMouth),     // 笑嘴巴
	types.MouthShape(NervousMouth),      // 紧张嘴巴
	types.MouthShape(PuckerMouth),       // 撅嘴
	types.MouthShape(SadMouth),          // 悲伤嘴巴
	types.MouthShape(SmileMouth),        // 微笑嘴巴
	types.MouthShape(SmirkMouth),        // 轻蔑嘴巴
	types.MouthShape(SurprisedMouth),    // 惊讶嘴巴
	types.BeardShape(ScruffBeard),       // 短胡子
	types.BeardShape(NoneBeard),         // 无胡子
	types.GlassesShape(RoundGlasses),    // 圆形眼镜
	types.GlassesShape(SquareGlasses),   // 方形眼镜
	types.GlassesShape(NoneGlasses),     // 无眼镜
	types.ClothesShape(CrewClothes),     // 圆领衣服
	types.ClothesShape(CollaredClothes), // 有领衣服
	types.ClothesShape(OpenClothes),     // 开放式衣服
}
