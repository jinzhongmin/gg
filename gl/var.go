package gl

import "github.com/jinzhongmin/gg/cc"

func init() { cc.Open("Opengl32.dll") }

var (
	glClearColor = cc.DlFunc[func(R, G, B, A float32), cc.Void]{Name: "glClearColor"}
	glClear      = cc.DlFunc[func(uint32), cc.Void]{Name: "glClear"}
)
