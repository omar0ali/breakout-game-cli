package core

import "github.com/omar0ali/breakout-game-cli/utils"

type Debug struct {
	message []string
	point   utils.Point
	config  *utils.Config
	visible bool
}

func CreateDebug(startPoint utils.Point, cfg *utils.Config) Debug {
	return Debug{
		message: []string{},
		point:   startPoint,
		config:  cfg,
	}
}

func (d *Debug) String() []string {
	return d.message
}

func (d *Debug) AddLine(line string) {
	d.message = append(d.message, line)
}

func (d *Debug) SetVisibility(vis bool) {
	d.visible = vis
}

func (d *Debug) Draw(window *Window) {
	if !d.config.Core.Debug {
		return
	}
	for y, item := range d.message {
		for x, letter := range item {
			window.SetContent(int(d.point.X)+x, int(d.point.Y)+y, letter)
		}
	}
	d.message = nil
}
