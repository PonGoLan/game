package pong

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type config struct {
	Scale        float64
	WindowTitle  string
	WindoWConfig pixelgl.WindowConfig
}

var c *config

func init() {
	c = new(config)

	c.Scale = 10
	c.WindowTitle = "PonGoLan"
	c.WindoWConfig = pixelgl.WindowConfig{
		Title:  c.WindowTitle,
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
}

func GetConfig() *config {
	return c
}

func (ci *config) GetTitle(frames int) string {
	return fmt.Sprintf("%s | FPS: %d", ci.WindowTitle, frames)
}
