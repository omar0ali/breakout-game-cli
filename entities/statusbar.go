package entities

import "fmt"

type Section int

const (
	TOPRIGHT Section = iota
	BOTTOMRIGHT
	BOTTOMLEFT
)

type SectionData struct {
	Lines []string
}

type StatusBar struct {
	Sections map[Section]*SectionData
	Ctx      *GameContext
}

func CreateStatusBar(ctx *GameContext) *StatusBar {
	return &StatusBar{
		Sections: map[Section]*SectionData{
			TOPRIGHT:    {Lines: []string{}},
			BOTTOMRIGHT: {Lines: []string{}},
			BOTTOMLEFT:  {Lines: []string{}},
		}, Ctx: ctx,
	}
}

func (s *StatusBar) Update(ctx *GameContext, dt float64) {
	s.AddLine(fmt.Sprintf("Balls: %v", ctx.Player.Balls), TOPRIGHT)
	s.AddLine(fmt.Sprintf("Bricks: %v", totalBricks), TOPRIGHT)
	// s.AddLine(fmt.Sprintf("Balls: %v", ctx.Player.Balls), BOTTOMRIGHT)
	// s.AddLine(fmt.Sprintf("Balls: %v", ctx.Player.Balls), BOTTOMLEFT)
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
	}
}

func (s *StatusBar) topRight(str string) {
	s.Sections[TOPRIGHT].Lines = append(s.Sections[TOPRIGHT].Lines, str)
}

func (s *StatusBar) bottomRight(str string) {
	s.Sections[BOTTOMRIGHT].Lines = append(s.Sections[BOTTOMRIGHT].Lines, str)
}

func (s *StatusBar) bottomLeft(str string) {
	s.Sections[BOTTOMLEFT].Lines = append(s.Sections[BOTTOMLEFT].Lines, str)
}
