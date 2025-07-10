package main

import (
	"time"

	"github.com/gdamore/tcell/v2"
)

type Window struct {
	ScreenWidth  int
	ScreenHeight int
	ticker       *time.Ticker
	screen       tcell.Screen
	style        tcell.Style
}

func CreateWindow(title string, frames time.Duration) (*Window, error) {
	screen, err := tcell.NewScreen()
	if err != nil {
		return nil, err
	}
	if err = screen.Init(); err != nil {
		return nil, err
	}
	width, height := screen.Size()
	screen.SetTitle(title)
	screen.Clear()

	window := &Window{
		ScreenWidth:  width,
		ScreenHeight: height,
		screen:       screen,
		ticker:       time.NewTicker(frames * time.Millisecond),
		style:        tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorGreenYellow),
	}
	return window, nil
}

func (s *Window) InitEventsKeys(callbackEvents func(*tcell.EventKey), callbackFrames func(), exit chan int) {
	go func() {
		for {
			events := s.screen.PollEvent()
			switch ev := events.(type) {
			case *tcell.EventResize:
				s.screen.Sync()
			case *tcell.EventKey:
				if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
					s.Close()
					exit <- 0
					return
				}
				callbackEvents(ev)
			}
		}
	}()
	go func() {
		for range s.ticker.C {
			s.screen.Clear()
			callbackFrames()
			s.screen.Show()
		}
	}()
}

func (s *Window) GetScreenSize() (int, int) {
	s.screen.Sync()
	s.ScreenWidth, s.ScreenHeight = s.screen.Size()
	return s.ScreenWidth, s.ScreenHeight
}

func (s *Window) Close() {
	s.screen.Fini()
	s.ticker.Stop()
}

func (s *Window) SetContent(x, y int, prune rune) {
	s.screen.SetContent(x, y, prune, nil, s.style)
}
