package main

import (
	"math"

	"github.com/ironverse/core"
)

var (
	SkyIslandsFrequency = 0.025
	PillarsFrequency    = 0.25
	HillsFrequency      = 0.025
	PlainsFrequency     = 0.01

	SkyIslandsHeight = 40
	PillarsHeight    = 15
	HillsHeight      = 25
	PlainsHeight     = 7
)

func main() {
	core.WorldGen.SetGenerator(GenerateWorld)
}
func GenerateWorld(data []byte, chunkRadius int, chunk *core.Position, hexOffset *core.Position) {
	i := 0
	for x := -chunkRadius; x <= chunkRadius; x++ {
		for z := -chunkRadius; z <= chunkRadius; z++ {
			if math.Abs(float64(x-z)) <= float64(chunkRadius) {
				for y := -chunkRadius; y <= chunkRadius; y++ {

					skyIslandValue := core.WorldGen.Get2dNoise(float64(hexOffset.X+x)*SkyIslandsFrequency+10000, float64(hexOffset.Z+z)*SkyIslandsFrequency+10000)
					pillarsValue := core.WorldGen.Get2dNoise(float64(hexOffset.X+x)*PillarsFrequency-10000, float64(hexOffset.Z+z)*PillarsFrequency-10000)
					hillsValue := core.WorldGen.Get2dNoise(float64(hexOffset.X+x)*HillsFrequency+10000, float64(hexOffset.Z+z)*HillsFrequency-10000)
					plainsValue := core.WorldGen.Get2dNoise(float64(hexOffset.X+x)*HillsFrequency+10000, float64(hexOffset.Z+z)*HillsFrequency-10000)

					if hexOffset.Y+y > 0 && hexOffset.Y+y < SkyIslandsHeight && skyIslandValue > 0.25 {
						//sky islands
					} else if pillarsValue > 0.25 {
						//pillars
						height := pillarsValue * PillarsHeight
						if height < 25 {
							height *= 0.25
						}
						if height >= 25 {
							height *= 1.25
						}
						if height > PillarsHeight {
							height = PillarsHeight
						}
						if hexOffset.Y+y+50 <= int(height) {
							data[i] = 1
						} else {
							data[i] = 0
						}
					} else if hillsValue > 0.25 {
						//hills
						height := hillsValue * HillsHeight
						if hexOffset.Y+y+50 <= int(height) {
							data[i] = 1
						} else {
							data[i] = 0
						}
					} else {
						//plains
						height := plainsValue * PlainsHeight
						if hexOffset.Y+y+50 <= int(height) {
							data[i] = 1
						} else {
							data[i] = 0
						}
					}
					i++
				}
			}
		}
	}
}
