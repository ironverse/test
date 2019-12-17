
Iron.V0.SetTerrainGenerator("terrain_generator")
function terrain_generator (seed, chunkRadius, data, offsetX, offsetY, offsetZ) {
    noise.seed(seed)

    var frequency = 0.025
    var heightRange = 50.0

    var i = 0
    for(var x = -chunkRadius; x <= chunkRadius; x++) {
        for(var z = -chunkRadius; z <= chunkRadius; z++) {
            if(Math.abs(x-z) <= chunkRadius) {
                for(var y = -chunkRadius; y <= chunkRadius; y++) {
                    
                    data[i] = 0
                    var value = 1.0 //noise.simplex2((offsetX+x) * frequency, (offsetZ+z) * frequency)
                    var height = (value + 1) * 0.5 * heightRange
                    if(offsetY+y <= height) {
                        data[i] = 1
                    }

                    i++
                }
            }
        }
    }
}