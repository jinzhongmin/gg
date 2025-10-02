package gl

func ClearColor(R, G, B, A float32) { glClearColor.Fn()(R, G, B, A) }
func Clear(bufferBit uint32)        { glClear.Fn()(bufferBit) }
