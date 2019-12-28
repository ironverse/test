package main

import (
	"math"

	"github.com/ironverse/core"
)

var (
	BiomesFrequency = 0.002

	SkyIslandsFrequency        = 0.025
	SkyIslandsSpacingFrequency = 0.25
	PillarsFrequency           = 0.25
	HillsFrequency             = 0.025
	PlainsFrequency            = 0.01

	SkyIslandsHeight = 40
	PillarsHeight    = 40
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

					biomesValue := core.WorldGen.Get2dNoise(float64(hexOffset.X+x)*BiomesFrequency, float64(hexOffset.Z+z)*BiomesFrequency)
					biomesTypeValue := core.WorldGen.Get2dNoise(float64(hexOffset.X+x)*BiomesFrequency+5000, float64(hexOffset.Z+z)*BiomesFrequency+5000)

					if biomesValue > 0 && biomesTypeValue > 0 {
						//sky islands
						if hexOffset.Y+y > 0 && hexOffset.Y+y < SkyIslandsHeight {
							skyIslandValue := core.WorldGen.Get2dNoise(float64(hexOffset.X+x)*SkyIslandsFrequency+10000, float64(hexOffset.Z+z)*SkyIslandsFrequency+10000)
							height := skyIslandValue * SkyIslandsHeight
							if hexOffset.Y+y <= int(height) {
								data[i] = 1
							} else {
								data[i] = 0
							}

							caveX := float64(hexOffset.X+x) * SkyIslandsSpacingFrequency
							caveY := float64(hexOffset.Y+y) * SkyIslandsSpacingFrequency
							caveZ := float64(hexOffset.Z+z) * SkyIslandsSpacingFrequency
							value := core.WorldGen.Get3dNoise(caveX, caveY, caveZ)
							if value > -0.5 {
								data[i] = 0
							}
						}
					} else if biomesValue <= 0 && biomesTypeValue <= 0 {
						//pillars
						pillarsValue := core.WorldGen.Get2dNoise(float64(hexOffset.X+x)*PillarsFrequency-10000, float64(hexOffset.Z+z)*PillarsFrequency-10000)
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
						if hexOffset.Y+y <= int(height) {
							data[i] = 1
						} else {
							data[i] = 0
						}
					} else if biomesValue > 0 && biomesTypeValue <= 0 {
						//hills
						hillsValue := core.WorldGen.Get2dNoise(float64(hexOffset.X+x)*HillsFrequency+10000, float64(hexOffset.Z+z)*HillsFrequency-10000)
						height := hillsValue * HillsHeight * math.Abs(biomesValue) * math.Abs(biomesTypeValue)
						if hexOffset.Y+y <= int(height) {
							data[i] = 1
						} else {
							data[i] = 0
						}
					} else if biomesValue <= 0 && biomesTypeValue > 0 {
						//plains
						plainsValue := core.WorldGen.Get2dNoise(float64(hexOffset.X+x)*HillsFrequency-10000, float64(hexOffset.Z+z)*HillsFrequency+10000)
						height := plainsValue * PlainsHeight * math.Abs(biomesValue) * math.Abs(biomesTypeValue)
						if hexOffset.Y+y <= int(height) {
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
