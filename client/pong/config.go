package pong

type config struct {
	Scale float64
}

var c *config

func init() {
	c = new(config)

	c.Scale = 10
}

func GetConfig() *config {
	return c
}
