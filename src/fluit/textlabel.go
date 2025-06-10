package fluit

import (
	"WormsWasher/src/assets"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type TextLabel struct {
	Instance

	size     Vector2[int]
	position Vector2[float64]
	color    color.Color
	opacity  float64

	fontSize float64
	AlignX   text.Align
	AlignY   text.Align
	fontName string
	text     string

	parent Instance

	Click      func(*TextLabel)
	MouseEnter func(*TextLabel)
	MouseLeave func(*TextLabel)
}

func (e *TextLabel) Update(dt float32) {
	mouseX, mouseY := ebiten.CursorPosition()

	if mouseX >= int(e.position.X) && mouseX <= int(e.position.X)*e.size.X && mouseY >= int(e.position.Y) && mouseY <= int(e.position.Y)*e.size.Y {
		if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
			if e.Click != nil {
				e.Click(e)
			}
		}
		if e.MouseEnter != nil {
			e.MouseEnter(e)
		}
	} else {
		if e.MouseLeave != nil {
			e.MouseLeave(e)
		}
	}
}

func (e TextLabel) Draw(screen *ebiten.Image) {
	op := &text.DrawOptions{}
	op.GeoM.Translate(e.position.X, e.position.Y)
	op.ColorScale.ScaleWithColor(e.color)
	op.ColorScale.ScaleAlpha(float32(e.opacity))
	op.PrimaryAlign = e.AlignX
	op.SecondaryAlign = e.AlignY

	text.Draw(screen, e.text, &text.GoTextFace{
		Source: assets.Fonts[e.fontName],
		Size:   e.fontSize,
	}, op)
}

func NewTextLabel() TextLabel {
	return TextLabel{
		opacity:  1,
		fontSize: 16,
	}
}

func (e *TextLabel) SetFontSize(v float64) *TextLabel {
	e.fontSize = v
	return e
}

func (e *TextLabel) SetAlign(x text.Align, y text.Align) *TextLabel {
	e.AlignX = x
	e.AlignY = y
	return e
}

func (e *TextLabel) SetFontName(v string) *TextLabel {
	e.fontName = v
	return e
}

func (e *TextLabel) SetText(v string) *TextLabel {
	e.text = v
	return e
}

func (e *TextLabel) SetSize(v Vector2[int]) *TextLabel {
	e.size = v
	return e
}
func (e *TextLabel) SetPosition(v Vector2[float64]) *TextLabel {
	e.position = v
	return e
}
func (e *TextLabel) SetColor(v color.Color) *TextLabel {
	e.color = v
	return e
}
func (e *TextLabel) SetOpacity(v float64) *TextLabel {
	e.opacity = v
	return e
}

func (e TextLabel) GetFontSize() float64 {
	return e.fontSize
}
func (e TextLabel) GetFontName() string {
	return e.fontName
}
func (e TextLabel) GetText() string {
	return e.text
}

func (e TextLabel) GetSize() Vector2[int] {
	return e.size
}
func (e TextLabel) GetPosition() Vector2[float64] {
	return e.position
}
func (e TextLabel) GetColor() color.Color {
	return e.color
}
func (e TextLabel) GetOpacity() float64 {
	return e.opacity
}
func (e *TextLabel) SetParent(v Instance) {
	e.parent = v
}
func (e *TextLabel) GetParent() Instance {
	return e.parent
}
