// Package core
package core

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/omar0ali/breakout-game-cli/utils"
)

type Window struct {
	Ticker *time.Ticker
	Screen tcell.Screen
	Style  tcell.Style
}

var delta float64

func CreateWindow(title string, cfg *utils.Config) (*Window, error) {
	screen, err := tcell.NewScreen()
	if err != nil {
		return nil, err
	}
	if err = screen.Init(); err != nil {
		return nil, err
	}

	screen.SetTitle(title)
	screen.Clear()

	window := &Window{
		Screen: screen,
		Ticker: time.NewTicker(time.Duration(cfg.Core.DurationTicker) * time.Millisecond),
		Style:  tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorGreenYellow),
	}
	return window, nil
}

func (s *Window) InitEventsKeys(
	callbackEvents func(ev *tcell.EventKey, delta float64),
	callbackFrames func(delta float64),
	exit chan int,
) {
	go func() {
		for {
			events := s.Screen.PollEvent()
			switch ev := events.(type) {
			case *tcell.EventResize:
				s.Screen.Sync()
			case *tcell.EventKey:
				if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
					s.Close()
					exit <- 0
					return
				}
				callbackEvents(ev, delta)
			}
		}
	}()
	go func() {
		last := time.Now()
		for range s.Ticker.C {
			now := time.Now()
			delta = now.Sub(last).Seconds()
			last = now

			s.Screen.Clear()
			lenStr := []rune(fmt.Sprintf("Delta: %.4f s | FPS: %.2f", delta, (1 / delta)))
			for i, r := range lenStr {
				s.SetContent(i, 0, r)
			}
			callbackFrames(delta)
			s.Screen.Show()
		}
	}()
}

func (s *Window) GetScreenSize() (int, int) {
	s.Screen.Sync()
	return s.Screen.Size()
}

func (s *Window) Close() {
	s.Screen.Fini()
	s.Ticker.Stop()
}

func (s *Window) SetContent(x, y int, prune rune) {
	s.Screen.SetContent(x, y, prune, nil, s.Style)
}
