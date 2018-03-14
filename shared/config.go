package pong

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type config struct {
	Scale        float64
	WindoWConfig pixelgl.WindowConfig
}

var c *config

func init() {
	c = new(config)

	c.Scale = 10
	c.WindoWConfig = pixelgl.WindowConfig{
		Title:  "PonGoLan - Client",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
}

func GetConfig() *config {
	return c
}
