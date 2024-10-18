package external

import (
	"math/rand"

	"github.com/kamalyes/go-identicon/constants"
	"github.com/kamalyes/go-identicon/enums"
	"github.com/kamalyes/go-identicon/models"
	"github.com/kamalyes/go-identicon/types"
)

// getRandomValue 从数组中随机获取一个值
func getRandomValue[T comparable](arr []T, avoid []T, usually []T) T {
	// 过滤掉避免的值
	avoidMap := make(map[T]struct{})
	for _, v := range avoid {
		avoidMap[v] = struct{}{}
	}

	var filteredArr []T
	for _, item := range arr {
		if _, found := avoidMap[item]; !found {
			filteredArr = append(filteredArr, item)
		}
	}

	// 添加常用值
	for _, item := range usually {
		for i := 0; i < 15; i++ {
			filteredArr = append(filteredArr, item)
		}
	}

	randomIdx := rand.Intn(len(filteredArr))
	return filteredArr[randomIdx]
}

// getRandomFillColor 随机获取填充颜色
func getRandomFillColor(colors []types.CommonColors) types.CommonColors {
	return colors[rand.Intn(len(colors))]
}

// GetRandomAvatarOption 随机生成头像选项
func GetRandomAvatarOption(useOption types.AvatarOption) types.AvatarOption {
	gender := getRandomValue(models.AvatarSettings.Genders, []types.Gender{useOption.Gender}, nil)

	// 初始化顶部形状列表
	topList := []types.TopsShape{enums.DannyTops, enums.WaveTops, enums.PixieTops}

	// 如果性别是男性，添加额外的顶部形状
	if gender == enums.MaleGender {
		for _, shape := range models.AvatarSettings.TopsShapes {
			if !contains(topList, shape) {
				topList = append(topList, shape)
			}
		}
	}

	// 获取随机值
	avatarOption := types.AvatarOption{
		Gender:       gender,
		WrapperShape: getRandomValue(models.AvatarSettings.WrapperShapes, []types.WrapperShape{useOption.WrapperShape}, nil),
		Background: types.Background{
			BackgroundColor: getRandomValue(models.AvatarSettings.BackgroundColors,
				[]types.BackgroundColor{useOption.Background.BackgroundColor}, nil),
			BorderColor: getRandomValue(models.AvatarSettings.BorderColors,
				[]types.BorderColor{useOption.Background.BorderColor},
				[]types.BorderColor{types.BorderColor(constants.ColorTransparent)}),
		},
	}

	// 如果性别是男性，设置胡子形状
	if gender == enums.MaleGender {
		avatarOption.Widgets.Beard.Shape = getRandomValue(models.AvatarSettings.BeardShapes, nil, nil)
	}

	// 设置其他部件的形状和颜色
	avatarOption.Widgets.Face.Shape = getRandomValue(models.AvatarSettings.FaceShapes, nil, nil)
	avatarOption.Widgets.Face.FillColor = getRandomFillColor(models.AvatarSettings.CommonColors)
	avatarOption.Widgets.Tops.Shape = getRandomValue(topList, nil, nil)
	avatarOption.Widgets.Tops.FillColor = getRandomFillColor(models.AvatarSettings.CommonColors)

	// 主要挂件
	avatarOption.Widgets.Earrings.Shape = getRandomValue(models.AvatarSettings.EarringsShapes, nil, []types.EarringsShape{enums.NoneEarrings})
	avatarOption.Widgets.Glasses.Shape = getRandomValue(models.AvatarSettings.GlassesShapes, nil, []types.GlassesShape{enums.NoneGlasses})
	avatarOption.Widgets.Clothes.FillColor = getRandomFillColor(models.AvatarSettings.CommonColors)

	avatarOption.Widgets.Ear.Shape = getRandomValue(models.AvatarSettings.EarShapes, []types.EarShape{useOption.Widgets.Ear.Shape}, nil)
	avatarOption.Widgets.Eyebrows.Shape = getRandomValue(models.AvatarSettings.EyebrowsShapes, []types.EyebrowsShape{useOption.Widgets.Eyebrows.Shape}, nil)
	avatarOption.Widgets.Eyes.Shape = getRandomValue(models.AvatarSettings.EyesShapes, []types.EyesShape{useOption.Widgets.Eyes.Shape}, nil)
	avatarOption.Widgets.Nose.Shape = getRandomValue(models.AvatarSettings.NoseShapes, []types.NoseShape{useOption.Widgets.Nose.Shape}, nil)
	avatarOption.Widgets.Mouth.Shape = getRandomValue(models.AvatarSettings.MouthShapes, []types.MouthShape{useOption.Widgets.Mouth.Shape}, nil)
	avatarOption.Widgets.Clothes.Shape = getRandomValue(models.AvatarSettings.ClothesShapes, []types.ClothesShape{useOption.Widgets.Clothes.Shape}, nil)

	return avatarOption
}

// contains 检查切片中是否包含某个元素
func contains[T comparable](slice []T, item T) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
