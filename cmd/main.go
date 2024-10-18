package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/kamalyes/go-identicon/enums"
	"github.com/kamalyes/go-identicon/external"
	"github.com/kamalyes/go-identicon/types"
)

func main() {
	// 示例设置
	settings := types.AvatarOption{
		Gender:       enums.MaleGender,
		WrapperShape: enums.CircleWrapper,
		Background: types.Background{
			BackgroundColor: "#ffffff",
			BorderColor:     "#cccccc",
		},
		Widgets: struct {
			Beard struct {
				Shape     types.BeardShape
				FillColor types.CommonColors
			}
			Face struct {
				Shape     types.FaceShape
				FillColor types.CommonColors
			}
			Tops struct {
				Shape     types.TopsShape
				FillColor types.CommonColors
			}
			Ear struct {
				Shape types.EarShape
			}
			Earrings struct {
				Shape types.EarringsShape
			}
			Eyebrows struct {
				Shape types.EyebrowsShape
			}
			Eyes struct {
				Shape types.EyesShape
			}
			Nose struct {
				Shape types.NoseShape
			}
			Glasses struct {
				Shape types.GlassesShape
			}
			Mouth struct {
				Shape types.MouthShape
			}
			Clothes struct {
				Shape     types.ClothesShape
				FillColor types.CommonColors
			}
		}{
			Beard: struct {
				Shape     types.BeardShape
				FillColor types.CommonColors
			}{
				Shape:     "Goatee",
				FillColor: "#000000",
			},
			Face: struct {
				Shape     types.FaceShape
				FillColor types.CommonColors
			}{
				Shape:     "Round",
				FillColor: "#f1c27d",
			},
			Tops: struct {
				Shape     types.TopsShape
				FillColor types.CommonColors
			}{
				Shape:     "Danny",
				FillColor: "#ff0000",
			},
			Ear: struct {
				Shape types.EarShape
			}{
				Shape: "Small",
			},
			Earrings: struct {
				Shape types.EarringsShape
			}{
				Shape: "Hoop",
			},
			Eyebrows: struct {
				Shape types.EyebrowsShape
			}{
				Shape: "Straight",
			},
			Eyes: struct {
				Shape types.EyesShape
			}{
				Shape: "Round",
			},
			Nose: struct {
				Shape types.NoseShape
			}{
				Shape: "Small",
			},
			Glasses: struct {
				Shape types.GlassesShape
			}{
				Shape: "Round",
			},
			Mouth: struct {
				Shape types.MouthShape
			}{
				Shape: "Smile",
			},
			Clothes: struct {
				Shape     types.ClothesShape
				FillColor types.CommonColors
			}{
				Shape:     "Shirt",
				FillColor: "#00ff00",
			},
		},
	}

	// 随机生成头像选项
	avatarOption := external.GetRandomAvatarOption(settings)
	// 将 AvatarOption 转换为 JSON
	avatarOptionJSON, err := json.MarshalIndent(avatarOption, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling to JSON: %v", err)
	}

	// 打印 JSON
	fmt.Println(string(avatarOptionJSON))
}
