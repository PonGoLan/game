package pong

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type camera struct {
	Pos       pixel.Vec
	Speed     float64
	Zoom      float64
	ZoomSpeed float64
}

var instance *camera

func init() {
	instance = new(camera)

	instance.Pos = pixel.V(250, 250)
	instance.Speed = 500.0
	instance.Zoom = 1.0
	instance.ZoomSpeed = 1.2
}

func GetCamera() *camera {
	return instance
}

func ApplyMatrixToWindow(win *pixelgl.Window) {
	cam := pixel.IM.Scaled(
		instance.Pos,
		instance.Zoom,
	).Moved(win.Bounds().Center().Sub(instance.Pos))

	win.SetMatrix(cam)
}
