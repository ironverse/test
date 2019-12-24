package main

import (
	"math"

	"github.com/ironverse/core"
)

var (
	Frequency     = 0.025
	HeightRange   = 50
	CaveFrequency = 0.01
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

					//--Start Hex Logic--//
					hexValue := false
					xFloat := float64(hexOffset.X+x) * Frequency
					zFloat := float64(hexOffset.Z+z) * Frequency
					value := core.WorldGen.Get2dNoise(xFloat, zFloat)
					height := (value + 1) * 0.5 * HeightRange
					if hexOffset.Y+y+50 <= int(height) {
						hexValue = true
					}
					//--End Hex Logic--//

					if hexValue {
						data[i] = 1
					} else {
						data[i] = 0
					}

					caveX := float64(hexOffset.X+x) * CaveFrequency
					caveY := float64(hexOffset.Y+y) * CaveFrequency
					caveZ := float64(hexOffset.Z+z) * CaveFrequency
					value = core.WorldGen.Get3dNoise(caveX, caveY, caveZ)
					if value > 0.5 {
						data[i] = 0
					}

					i++
				}
			}
		}
	}
}
