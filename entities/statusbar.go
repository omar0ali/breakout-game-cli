package entities

import (
	"fmt"
	"time"
)

type Section int

const (
	TOPRIGHT Section = iota
	BOTTOMRIGHT
	BOTTOMLEFT
	CENTERSCREEN
)

var elapsedSeconds, elapsedMinutes int

type SectionData struct {
	Lines []string
}

type StatusBar struct {
	Sections  map[Section]*SectionData
	StartTime time.Time
	Ctx       *GameContext
}

func CreateStatusBar(ctx *GameContext) *StatusBar {
	starTime := time.Now()
	return &StatusBar{
		Sections: map[Section]*SectionData{
			TOPRIGHT:     {Lines: []string{}},
			BOTTOMRIGHT:  {Lines: []string{}},
			BOTTOMLEFT:   {Lines: []string{}},
			CENTERSCREEN: {Lines: []string{}},
		}, Ctx: ctx, StartTime: starTime,
	}
}

func (s *StatusBar) Update(ctx *GameContext, dt float64) {
	s.AddLine(fmt.Sprintf("Balls: %v", ctx.Player.Balls), TOPRIGHT)
	s.AddLine(fmt.Sprintf("Bricks: %v", totalBricks), TOPRIGHT)

	if totalBricks <= 0 {
		s.AddLine("------------", CENTERSCREEN)
		s.AddLine("| You Win! |", CENTERSCREEN)
		s.AddLine("------------", CENTERSCREEN)
		s.AddLine(fmt.Sprintf("%02d:%02d s",
			elapsedMinutes,
			elapsedSeconds,
		), CENTERSCREEN)
		return
	}

	if len(ctx.Balls) == 0 && ctx.Player.Balls == 0 {
		s.AddLine(fmt.Sprintln("GAME OVER"), BOTTOMRIGHT)
		s.AddLine(fmt.Sprintln("GAME OVER -\tPress q to quit."), BOTTOMLEFT)
		s.AddLine(fmt.Sprintf("%02d:%02d s",
			elapsedMinutes,
			elapsedSeconds,
		), BOTTOMRIGHT)
		return
	}
	if ballCounter == 0 {
		s.AddLine("-------------------------------------------------", CENTERSCREEN)
		s.AddLine("|  Click the left mouse button to get started!  |", CENTERSCREEN)
		s.AddLine("-------------------------------------------------", CENTERSCREEN)
		s.StartTime = time.Now()
		return
	}
	elapsed := time.Since(s.StartTime)
	elapsedMinutes = int(elapsed.Minutes())
	elapsedSeconds = int(elapsed.Seconds()) % 60
	s.AddLine(fmt.Sprintf("%02d:%02d s",
		elapsedMinutes,
		elapsedSeconds,
	), TOPRIGHT)
}

func (s *StatusBar) Draw(ctx *GameContext) {
	width, height := ctx.Window.GetScreenSize()
	for y, line := range s.Sections[TOPRIGHT].Lines {
		startPos := width - len(line)
		for index, j := range line {
			ctx.Window.SetContent(startPos+index, y, j)
		}
	}
	for i, line := range s.Sections[BOTTOMRIGHT].Lines {
		startPos := width - len(line)
		y := height - len(s.Sections[BOTTOMRIGHT].Lines) + i
		for index, ch := range line {
			ctx.Window.SetContent(startPos+index, y, ch)
		}
	}
	for i, line := range s.Sections[BOTTOMLEFT].Lines {
		y := height - len(s.Sections[BOTTOMLEFT].Lines) + i
		for index, ch := range line {
			ctx.Window.SetContent(index, y, ch)
		}
	}

	for i, line := range s.Sections[CENTERSCREEN].Lines {
		y := (height / 2) - len(s.Sections[CENTERSCREEN].Lines) + i
		halfWidth := width / 2
		halfWidth -= len(line) / 2
		for index, ch := range line {
			ctx.Window.SetContent(halfWidth+index, y, ch)
		}
	}

	for _, section := range s.Sections {
		section.Lines = section.Lines[:0]
	}
}

func (s *StatusBar) AddLine(str string, sec Section) {
	switch sec {
	case TOPRIGHT:
		s.topRight(str)
	case BOTTOMLEFT:
		s.bottomLeft(str)
	case BOTTOMRIGHT:
		s.bottomRight(str)
	case CENTERSCREEN:
		s.centerScreen(str)
	}
}

func (s *StatusBar) topRight(str string) {
	s.Sections[TOPRIGHT].Lines = append(s.Sections[TOPRIGHT].Lines, str)
}

func (s *StatusBar) centerScreen(str string) {
	s.Sections[CENTERSCREEN].Lines = append(s.Sections[CENTERSCREEN].Lines, str)
}

func (s *StatusBar) bottomRight(str string) {
	s.Sections[BOTTOMRIGHT].Lines = append(s.Sections[BOTTOMRIGHT].Lines, str)
}

func (s *StatusBar) bottomLeft(str string) {
	s.Sections[BOTTOMLEFT].Lines = append(s.Sections[BOTTOMLEFT].Lines, str)
}
